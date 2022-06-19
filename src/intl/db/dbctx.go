package db

import (
	"github.com/betalixt/eventSourceUsers/optn"
	"github.com/betalixt/eventSourceUsers/util/blerr"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDatabase(optn *optn.DatabaseOptions) *sqlx.DB {

	db, err := sqlx.Open("postgres", optn.ConnectionString)
	if err != nil {
		panic(blerr.NewError(
				blerr.DatabaseConnectionOpenFailure,
				500,
				err.Error(),
			))
	}

	err = db.Ping()
	if err != nil {
		panic(blerr.NewError(blerr.DatabasePingFailure, 500, err.Error()))
	}

	return db
}
