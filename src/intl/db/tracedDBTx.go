package db

import (
	"database/sql"
	"time"

	"github.com/betalixt/eventSourceUsers/intl/trace"
	"github.com/jmoiron/sqlx"
)

type TracedDBTransaction struct {
	*sqlx.Tx
	tracer      trace.ITracer
	serviceName string
}

func (tx *TracedDBTransaction) Get(
	dest interface{},
	query string,
	args ...interface{},
) error {
	start := time.Now()
	err := tx.Tx.Get(dest, query, args...)
	end := time.Now()
	if err != nil {
		tx.tracer.TraceDependency(
			"",
			tx.DriverName(),
			tx.serviceName,
			"Get",
			false,
			start,
			end,
			trace.NewField("error", err.Error()),
			trace.NewField("query", query),
		)
	} else {
		tx.tracer.TraceDependency(
			"",
			tx.DriverName(),
			tx.serviceName,
			"Get",
			true,
			start,
			end,
		)
	}
	return err
}

func (tx *TracedDBTransaction) Select(
	dest interface{},
	query string,
	args ...interface{},
) error {
	start := time.Now()
	err := tx.Tx.Select(dest, query, args...)
	end := time.Now()
	if err != nil {
		tx.tracer.TraceDependency(
			"",
			tx.DriverName(),
			tx.serviceName,
			"Get",
			false,
			start,
			end,
			trace.NewField("error", err.Error()),
			trace.NewField("query", query),
		)
	} else {
		tx.tracer.TraceDependency(
			"",
			tx.DriverName(),
			tx.serviceName,
			"Get",
			true,
			start,
			end,
		)
	}
	return err
}

func (tx *TracedDBTransaction) Exec(
	query string,
	args ...interface{},
) (sql.Result, error) {
	start := time.Now()
	res, err := tx.Tx.Exec(query, args...)
	end := time.Now()
	if err != nil {
		tx.tracer.TraceDependency(
			"",
			tx.DriverName(),
			tx.serviceName,
			"Get",
			false,
			start,
			end,
			trace.NewField("error", err.Error()),
			trace.NewField("query", query),
		)
	} else {
		tx.tracer.TraceDependency(
			"",
			tx.DriverName(),
			tx.serviceName,
			"Get",
			true,
			start,
			end,
		)
	}
	return res, err
}
