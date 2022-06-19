package db

import (
	"database/sql"
	"time"

	"github.com/betalixt/eventSourceUsers/intl/trace"
	"github.com/jmoiron/sqlx"
)

type TracedDBContext struct {
	*sqlx.DB
	tracer      trace.ITracer
	serviceName string
}

func NewTracedDBContext(
	db *sqlx.DB,
	tracer trace.ITracer,
	serviceName string,
) *TracedDBContext {

	return &TracedDBContext{
		tracer:      tracer,
		DB:          db,
		serviceName: serviceName,
	}
}

func (trDB *TracedDBContext) Get(
	dest interface{},
	query string,
	args ...interface{},
) error {
	start := time.Now()
	err := trDB.DB.Get(dest, query, args...)
	end := time.Now()
	if err != nil {
		trDB.tracer.TraceDependency(
			"",
			trDB.DriverName(),
			trDB.serviceName,
			"Get",
			false,
			start,
			end,
			trace.NewField("error", err.Error()),
			trace.NewField("query", query),
		)
	} else {
		trDB.tracer.TraceDependency(
			"",
			trDB.DriverName(),
			trDB.serviceName,
			"Get",
			true,
			start,
			end,
		)
	}
	return err
}

func (trDB *TracedDBContext) Select(
	dest interface{},
	query string,
	args ...interface{},
) error {

	start := time.Now()
	err := trDB.DB.Select(dest, query, args...)
	end := time.Now()
	if err != nil {
		trDB.tracer.TraceDependency(
			"",
			trDB.DriverName(),
			trDB.serviceName,
			"Get",
			false,
			start,
			end,
			trace.NewField("error", err.Error()),
			trace.NewField("query", query),
		)
	} else {
		trDB.tracer.TraceDependency(
			"",
			trDB.DriverName(),
			trDB.serviceName,
			"Get",
			true,
			start,
			end,
		)
	}
	return err
}

func (trDB *TracedDBContext) Exec(
	query string,
	args ...interface{},
) (sql.Result, error) {

	start := time.Now()
	res, err := trDB.DB.Exec(query, args...)
	end := time.Now()
	if err != nil {
		trDB.tracer.TraceDependency(
			"",
			trDB.DriverName(),
			trDB.serviceName,
			"Get",
			false,
			start,
			end,
			trace.NewField("error", err.Error()),
			trace.NewField("query", query),
		)
	} else {
		trDB.tracer.TraceDependency(
			"",
			trDB.DriverName(),
			trDB.serviceName,
			"Get",
			true,
			start,
			end,
		)
	}
	return res, err
}

func (db *TracedDBContext) Beginx() (*TracedDBTransaction, error) {
	tx, err := db.DB.Beginx()
	return &TracedDBTransaction{
		Tx:          tx,
		tracer:      db.tracer,
		serviceName: db.serviceName,
	}, err
}

func (db *TracedDBContext) MustBegin() *TracedDBTransaction {
	tx := db.DB.MustBegin()
	return &TracedDBTransaction{
		Tx:          tx,
		tracer:      db.tracer,
		serviceName: db.serviceName,
	}
}
