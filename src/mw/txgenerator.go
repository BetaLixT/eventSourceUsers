package mw

import (
	// "strings"

	"fmt"
	"strings"

	"github.com/betalixt/eventSourceUsers/intl/trace"
	"github.com/betalixt/eventSourceUsers/intl/trace/hlpr"
	"github.com/betalixt/eventSourceUsers/util/txcontext"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	// "github.com/betalixt/eventSourceUsers/util/blerr"
	"github.com/gin-gonic/gin"
)

func TransactionContextGenerationMiddleware(
	ins *trace.AppInsightsCore,
	lgr *zap.Logger,
	db *sqlx.DB,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trcprnt := ctx.GetHeader("traceparent")
		tid := ""
		pid := ""
		rid := ""
		flg := ""
		var err error

		// TODO benchmark and optimize
		if trcprnt != "" {
			// this is mostly to validate that the traceparent is legit
			var pidr []byte	
			_, _, pidr, _, err = hlpr.ParseTraceparentRaw(trcprnt)
			if err != nil {
				trcprnt, err = hlpr.GenerateNewTraceparent(true)
				if err != nil {
					lgr.Error("Failed to generate traceparent", zap.Error(err))
				} else {	
					values := strings.Split(trcprnt, "-")
					tid = values[1]
					rid = values[2]
					flg = values[3]
				}
			} else {
				values := strings.Split(trcprnt, "-")
				tid = values[1]
				flg = values[3]
				rid, err = hlpr.GenerateParentId()
				if err != nil {
					lgr.Error("Failed to generate parent id", zap.Error(err))
				} else {
					if err := hlpr.ValidateParentIdValue(pidr); err == nil {
						pid = values[2]
					} 
					
					trcprnt = fmt.Sprintf(
						"%s-%s-%s-%s",
						values[0],
						values[1],
						rid,
						values[3],
					)
				}
			}

		} else {
			trcprnt, err = hlpr.GenerateNewTraceparent(true)
				if err != nil {
					lgr.Error("Failed to generate traceparent", zap.Error(err))
				} else {	
					values := strings.Split(trcprnt, "-")
					tid = values[1]
					rid = values[2]
					flg = values[3]
				}
		} 
		

		tctx := txcontext.NewTransactionContext(
			trcprnt,
			tid,
			pid,
			rid,
			flg,
			db,
			ins,
			lgr,
		)
		ctx.Set("tctx", tctx)
		ctx.Writer.Header().Set("traceparent", trcprnt)
		ctx.Next()
	}
}
