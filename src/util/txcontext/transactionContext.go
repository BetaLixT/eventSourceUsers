package txcontext

import (
	"github.com/betalixt/eventSourceUsers/intl/db"
	"github.com/betalixt/eventSourceUsers/intl/http"
	"github.com/betalixt/eventSourceUsers/intl/trace"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type TransactionContext struct {
	traceparent     string
	tid             string
	pid             string
	rid             string
	flg             string
	isParent        bool
	db              *sqlx.DB
	logger          *zap.Logger
	appinsightsCore *trace.AppInsightsCore
	tracer          trace.ITracer
	httpClient      *http.HttpClient
	tracedDB        *db.TracedDBContext
}

func (tctx *TransactionContext) GetLogger() *zap.Logger {
	return tctx.logger
}
func (tctx *TransactionContext) GetHttpClient() *http.HttpClient {
	if tctx.httpClient == nil {
		tctx.httpClient = http.NewClient(
			tctx.GetTracer(),
			nil,
			tctx.tid,
			tctx.pid,
			tctx.flg,
		)
	}
	return tctx.httpClient
}
func (tctx *TransactionContext) GetDatabaseContext() *db.TracedDBContext {
	if tctx.tracedDB == nil {
		tctx.tracedDB = db.NewTracedDBContext(
			tctx.db,
			tctx.GetTracer(),
			"main-database",
		)
	}
	return tctx.tracedDB
}

func (tctx *TransactionContext) GetTracer() trace.ITracer {
	if tctx.tracer == nil {
		tctx.tracer = trace.NewAppInsightsTrace(
			tctx.appinsightsCore,
			tctx.tid,
			tctx.pid,
			tctx.rid,
		)
	}
	return tctx.tracer
}

func (tctx *TransactionContext) IsParent() bool {
	return tctx.isParent
}

// - Constructor
func NewTransactionContext(
	traceparent string,
	tid string,
	pid string,
	rid string,
	flg string,
	db *sqlx.DB,
	appinsightsCore *trace.AppInsightsCore,
	logger *zap.Logger,
) *TransactionContext {

	return &TransactionContext{
		tid:             tid,
		pid:             pid,
		rid:             rid,
		flg:             flg,
		traceparent:     traceparent,
		appinsightsCore: appinsightsCore,
		db:              db,
		logger:          logger.With(zap.String("tid", tid), zap.String("pid", pid), zap.String("rid", rid)),
	}
}
