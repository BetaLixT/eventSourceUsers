package mw

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/betalixt/eventSourceUsers/svc"
	"github.com/betalixt/eventSourceUsers/util/blerr"
	"github.com/betalixt/eventSourceUsers/util/txcontext"
)

func AuthMiddleware(tknSvc *svc.TokenService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		authSplit := strings.Split(authHeader, " ")
		if len(authSplit) != 2 || authSplit[0] != "Bearer" {
			ctx.Error(blerr.NewError(blerr.TokenInvalidCode, 401, ""))
			ctx.Abort()
			return
		}

		// Dependent on the txgenerator
		tctx := ctx.MustGet("tctx").(*txcontext.TransactionContext)
		_, err := tknSvc.ValidateToken(tctx, authSplit[1])
		if err != nil {
			ctx.Error(err)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
