package mw

import (
	"errors"

	"github.com/betalixt/eventSourceUsers/util/blerr"
	"github.com/betalixt/eventSourceUsers/util/txcontext"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tctx := ctx.MustGet("tctx").(*txcontext.TransactionContext)
		ctx.Next()
		lgr := tctx.GetLogger()

		if len(ctx.Errors) > 0 {
			errs := make([]error, len(ctx.Errors))
			berr := (*blerr.Error)(nil)
			var temp *blerr.Error
			for idx, err := range ctx.Errors {
				errs[idx] = err.Err
				if berr != nil && errors.As(err.Err, &temp) {
					berr = temp
				}
			}
			lgr.Error("errors processing request", zap.Errors("error", errs))
			if berr != nil {
				ctx.JSON(berr.StatusCode, berr)
			} else {
				ctx.JSON(500, blerr.UnexpectedError())
			}
		} else {
			if !ctx.Writer.Written() {
				lgr.Error("No response was written")
				ctx.JSON(500, blerr.UnsetResponseError())
			}
		}

	}
}
