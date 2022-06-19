package main

import (
	"github.com/betalixt/eventSourceUsers/ctrl"
	"github.com/betalixt/eventSourceUsers/intl/trace"
	"github.com/betalixt/eventSourceUsers/mw"
	"github.com/betalixt/eventSourceUsers/optn"
	"github.com/betalixt/eventSourceUsers/svc"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func NewGinEngine(
	lgr *zap.Logger,
	corsOptn *optn.CorsOptions,
	tknSvc *svc.TokenService,
	db *sqlx.DB,
	ins *trace.AppInsightsCore,
	attctrl *ctrl.AttachmentController,
) *gin.Engine {
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	router.SetTrustedProxies(nil)

	// - Setting up middlewares
	router.Use(mw.TransactionContextGenerationMiddleware(ins, lgr, db))
	router.Use(mw.LoggingMiddleware())
	router.Use(mw.RecoveryMiddleware(lgr))
	router.Use(mw.ErrorHandlerMiddleware())
	router.Use(mw.CorsMiddleware(lgr, corsOptn))
  // TODO Make this configurable
	// router.Use(mw.AuthMiddleware(tknSvc))

	// - Responding to head
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "alive",
		})
	})
	v1 := router.Group("api/v1")
	attctrl.RegisterRoutes(v1.Group("attachments"))

	return router
}
