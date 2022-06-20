package common

import "database/sql"

type IDatabaseContext interface {
	Beginx() (*IDatabaseTransaction, error)
}

type IDatabaseTransaction interface {
	Get(
		dest interface{},
		query string,
		args ...interface{},
	) error

	Select(
		dest interface{},
		query string,
		args ...interface{},
	) error

	Exec(
		query string,
		args ...interface{},
	) (sql.Result, error)

	Commit() error
	Rollback() error
}
