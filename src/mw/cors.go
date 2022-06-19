package mw

import (
	"github.com/betalixt/eventSourceUsers/optn"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CorsMiddleware(lgr *zap.Logger, optn *optn.CorsOptions) gin.HandlerFunc {
	corsCfg := cors.DefaultConfig()
	corsCfg.AllowOrigins = optn.AllowedOrigins
	corsCfg.AllowCredentials = true
	corsCfg.AllowHeaders = []string {
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"X-CSRF-Token",
		"Authorization",
		"accept",
		"origin",
		"Cache-Control",
		"X-Requested-With",
	}
	lgr.Info(
		"Configuring cors",
		zap.Strings("allowedOrigins", corsCfg.AllowOrigins),
	)
	return cors.New(corsCfg)
}
