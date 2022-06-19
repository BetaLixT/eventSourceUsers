package mw

import (
	"errors"
	"net"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/betalixt/eventSourceUsers/util/blerr"
	"github.com/betalixt/eventSourceUsers/util/txcontext"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RecoveryMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		tctxAny, exists := c.Get("tctx") 
		defer func() {
			if err := recover(); err != nil {
				
				// Dependent on the txgenerator
				// Making it more resilient to avoid errors
				lgr := (*zap.Logger)(nil)
				if !exists {
					lgr = logger
				} else if tctx, ok := tctxAny.(*txcontext.TransactionContext); !ok {
					lgr = logger
				} else {
					lgr = tctx.GetLogger()
				}

				// In case the get logger fails
				if lgr == nil {
					lgr = logger
				}	
				
				perr, ok := err.(blerr.Error)
				if ok {
					c.JSON(perr.StatusCode, perr)
				} else {
					// Check for a broken connection, as it is not really a
					// condition that warrants a panic stack trace.
					var brokenPipe bool
					if ne, ok := err.(*net.OpError); ok {
						var se *os.SyscallError
						if errors.As(ne, &se) {
							if strings.Contains(
								strings.ToLower(se.Error()), "broken pipe") ||
								strings.Contains(strings.ToLower(se.Error()),
								"connection reset by peer",
								) {
								brokenPipe = true
							}
						}
					}
					
					httpRequest, _ := httputil.DumpRequest(c.Request, false)
					headers := strings.Split(string(httpRequest), "\r\n")
					for idx, header := range headers {
						current := strings.Split(header, ":")
						if current[0] == "Authorization" {
							headers[idx] = current[0] + ": *"
						}
					}
					headersToStr := strings.Join(headers, "\r\n")
					if brokenPipe {	
						lgr.Error(
							"Panic recovered, broken pipe",
							zap.String("headers", headersToStr),
							zap.Any("error", err),
						)
						c.Abort()
					} else {	
						lgr.Error(
							"Panic recovered",
							zap.String("headers", headersToStr),
							zap.Any("error", err),
							zap.Stack("stack"),
						)
						c.JSON(500, blerr.UnexpectedError())
					}
				}
			}
		}()
		c.Next()
	}
}
