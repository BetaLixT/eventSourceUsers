package repo

import (
	"github.com/betalixt/eventSourceUsers/intl/db"
	"go.uber.org/zap"
)

func RunMigrations(
  dbctx *db.TracedDBContext,
  lgr *zap.Logger,
) error {
  return db.RunMigrations(lgr, dbctx, migrations)
}

var migrations = []db.MigrationScript {
	
}
