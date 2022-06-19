package main

import (
	"fmt"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/betalixt/eventSourceUsers/clnt"
	"github.com/betalixt/eventSourceUsers/ctrl"
	"github.com/betalixt/eventSourceUsers/intl/db"
	"github.com/betalixt/eventSourceUsers/intl/trace"
	"github.com/betalixt/eventSourceUsers/optn"
	"github.com/betalixt/eventSourceUsers/repo"
	"github.com/betalixt/eventSourceUsers/svc"
	"github.com/betalixt/eventSourceUsers/util/blerr"
	"github.com/betalixt/eventSourceUsers/util/config"
	"github.com/betalixt/eventSourceUsers/util/logger"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	// - Custom panic logging
	defer func() {
		if err := recover(); err != nil {	
			fmt.Printf("[Panic Encountered] %v\n", time.Now())
			prsd, ok := err.(blerr.Error); if ok {
				fmt.Printf(
					"ErrorCode: %d [%s]\n",
					prsd.ErrorCode,
					prsd.ErrorMessage,
				)
				if prsd.ErrorDetail != "" {
					fmt.Printf("Details: %s\n", prsd.ErrorDetail)
				}
			} else {
				fmt.Printf("error: %v\n", err)
			}
			fmt.Printf("%s", string(debug.Stack()))
		}
	}()

	app := fx.New(
		fx.Provide(config.NewConfig),
		fx.Provide(optn.NewCorsOptions),
		fx.Provide(optn.NewDatabaseOptions),
		fx.Provide(optn.NewAppInsightsOptions),
		fx.Provide(optn.NewFileServiceClientOptions),
		fx.Provide(logger.NewLogger),	
		fx.Provide(trace.NewAppInsightsCore),
		fx.Provide(db.NewDatabase),
		fx.Provide(clnt.NewFileServiceClient),
		fx.Provide(svc.NewTokenService),
		fx.Provide(svc.NewAttachmentService),
		fx.Provide(ctrl.NewAttachmentController),
		fx.Provide(NewGinEngine),
		fx.Invoke(StartService),
	)
	app.Run();
}

func StartService(
	cfg *viper.Viper,
	lgr *zap.Logger,
	dbctx *sqlx.DB,
	gin *gin.Engine,
	appi *trace.AppInsightsCore,
) {
	defer appi.Close()	
	port := cfg.GetString("PORT")
	if port == "" {
		lgr.Warn("No port was specified, using 8080")
		port = "8080"
	} else if _, err := strconv.Atoi(port); err != nil {
		lgr.Error("Non numeric value was specified for port")
		panic(blerr.NewError(blerr.InvalidPortCode, 500, ""))
	}

	lgr.Info("Running migrations")
	err := repo.RunMigrations(
		db.NewTracedDBContext(dbctx, trace.NewZapTracer(lgr), "main-database"),
		lgr,
	)
	if err != nil {
		lgr.Warn("Failed running migrations", zap.Error(err))
	}
	
	lgr.Info("Starting service", zap.String("port", port))
	gin.Run(fmt.Sprintf(":%s", port))
}
