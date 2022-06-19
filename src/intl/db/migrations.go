package db

import (
	"sort"
	"time"

	"github.com/betalixt/eventSourceUsers/util/blerr"
	"go.uber.org/zap"
)


func RunMigrations(
	lgr *zap.Logger,
	db *TracedDBContext,
	migrations []MigrationScript,
) error {

	tx := db.MustBegin()
	chck := ExistsEntity{}

	// - creating timestamp procedures if requried
	err := tx.Get(&chck, CheckTimestampProceduresExist)
	if err != nil {
		lgr.Error(
			"Failed fetching procedure info",
			zap.Error(err),
		)
		panic(blerr.NewError(
			blerr.DatabaseMigrationProcedureCheckFailedCode,
			500,
			err.Error(),
		))
	}

	if !chck.Exists {
		lgr.Info("Creating timestamp procedures")
		tx.MustExec(timestampProcedures.up)
	}

	// - creating migration table if required
	err = tx.Get(&chck, CheckMigrationExists)
	if err != nil {
		lgr.Error(
			"Failed fetching migration info",
			zap.Error(err),
		)
		panic(blerr.NewError(
			blerr.DatabaseMigrationMigrationCheckFailedCode,
			500,
			err.Error(),
		))
	}
	var exMigrs []migrationEntity

	if !chck.Exists {
		lgr.Info("Creating migration table")
		tx.MustExec(migrationTable.up)
		exMigrs = []migrationEntity{}
	} else {
		lgr.Info("Fetching migration history")
		err = tx.Select(&exMigrs, GetAllMigrations)
		if err != nil {
			lgr.Error(
				"failed to fetch migrations",
				zap.Error(err),
			)
			panic(blerr.NewError(
				blerr.DatabaseMigrationMigrationFetchFailedCode,
				500,
				err.Error(),
			))
		}
	}
	sort.Slice(exMigrs, func(i, j int) bool {
		return exMigrs[i].Index < exMigrs[j].Index
	})

	exMigrsLen := len(exMigrs)
	
	for idx, migr := range migrations {
		if idx < exMigrsLen {
			if migr.key != exMigrs[idx].Key {
				panic(blerr.NewError(
					blerr.DatabaseMigrationMigrationHistoryMismatchCode,
					500,
					""))
			}
		} else {
			lgr.Info("Running migration", zap.String("migration", migr.key))
			tx.MustExec(migr.up)
			tx.MustExec(AddMigration, migr.key)
		}
	}
	return tx.Commit()
}

type MigrationScript struct {
	key  string
	up   string
	down string
}

type migrationEntity struct {
	Index           int        `db:"index"`
	Key             string     `db:"key"`
	DateTimeCreated *time.Time `db:"datetimecreated"`
}

var timestampProcedures = MigrationScript{
	up: `
		CREATE OR REPLACE FUNCTION trigger_set_datetimecreated()
		RETURNS TRIGGER AS $$
		BEGIN
			NEW.dateTimeCreated = NOW();
			RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;
		
		CREATE OR REPLACE FUNCTION trigger_set_datetimeupdated()
		RETURNS TRIGGER AS $$
		BEGIN
			NEW.dateTimeUpdated = NOW();
			RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;`,
	down: `
		DROP FUNCTION trigger_set_datetimeupdated();
		DROP FUNCTION trigger_set_datetimecreated();`,
}

var migrationTable = MigrationScript{
	up: `
		CREATE TABLE migrations (
			index SERIAL,
			key text PRIMARY KEY,
			dateTimeCreated timestamp with time zone NULL
		);
		
		CREATE TRIGGER set_migrations_datetimecreated
		BEFORE INSERT ON migrations
		FOR EACH ROW
		EXECUTE PROCEDURE trigger_set_datetimecreated();`,
	down: `
		DROP TRIGGER set_migrations_datetimecreated on migrations;
		DROP TABLE Migrations;`,
}

const (
	CheckTimestampProceduresExist = `
		SELECT EXISTS(
			SELECT * FROM (
				SELECT Count(p.proname) as count
				FROM pg_proc AS p
				JOIN pg_namespace n ON p.pronamespace = n.oid
				WHERE p.proname in (
					'trigger_set_datetimecreated', 
					'trigger_set_datetimeupdated'
					) 
					AND n.nspname = 'public'
			) as c
			WHERE c.count = 2
		) as exists`
	CheckMigrationExists = `
		SELECT EXISTS(
			SELECT * FROM pg_tables
			WHERE schemaname = 'public' AND tablename = 'migrations'
		) as exists`
	GetAllMigrations = `
		SELECT * FROM migrations`
	AddMigration = `
		INSERT INTO migrations (key) VALUES ($1)`
)
