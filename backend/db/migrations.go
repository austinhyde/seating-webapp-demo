package db

import (
	"context"
	"strings"

	"github.com/austinhyde/seating/util"
	"github.com/rs/zerolog"

	"github.com/pkg/errors"
)

const MIGRATION_TABLE = "migration"

var migrations = map[int]string{
	1: `
		CREATE TABLE ` + MIGRATION_TABLE + ` (
			version int PRIMARY KEY,
			ts timestamptz DEFAULT now()
		);
		CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
	`,
	2: `
		CREATE TABLE location (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			name text NOT NULL,
			location geometry(point, 4326),
		)
	`,
}

// GetCurrentMigrationVersion returns the last applied migration number
func GetCurrentMigrationVersion(ctx context.Context, log *zerolog.Logger, db Queryable) (int, error) {
	version := 0
	err := db.QueryRow(ctx, `
		SELECT coalesce(
			(SELECT max(version)
			FROM `+MIGRATION_TABLE+`),
			0
		);
	`).Scan(&version)
	if err != nil && strings.Contains(err.Error(), `relation "`+MIGRATION_TABLE+`" does not exist`) {
		version = 0
		log.Warn().Err(err).Msg("migration table missing; ignoring error and treating as version 0")
		err = nil
	}
	return version, err
}

// ApplyMigrations looks up the current migration number from the database, then executes newer migration SQL
func ApplyMigrations(ctx context.Context, log *zerolog.Logger, db Transactable) error {
	log.Info().Msg("applying migrations")

	version, err := GetCurrentMigrationVersion(ctx, log, db)
	if err != nil {
		return errors.Wrap(err, "Could not find current version number")
	}
	log.Info().Int("version", version).Msg("detected migration version")

	migration := version + 1
	for {
		sql, ok := migrations[migration]
		if !ok {
			// TODO: can't differentiate between no-more-migrations and missing-migration cases
			// the former is not an error, but the latter is.
			// could do it by checking to see if there are any indexes > `migration`
			return nil
		}

		err = ApplyMigration(ctx, log, db, migration, sql)
		if err != nil {
			return errors.Wrapf(err, "Could not apply migration %d ('%s')", migration, util.Abbrev(sql, 20))
		}

		migration += 1
	}
}

// ApplyMigration applies a single `sql` migration (identified by `version`) in a transaction
func ApplyMigration(ctx context.Context, log *zerolog.Logger, db Transactable, version int, sql string) error {
	abbrev := util.Abbrev(sql, 20)
	log.Info().Int("version", version).Str("sql", abbrev).Msg("attempting to apply migration")

	return InTransaction(ctx, log, db, func(q Queryable) error {
		_, err := db.Exec(ctx, sql)
		if err != nil {
			return errors.Wrapf(err, "could not execute migration sql for migration %d ('%s')", version, abbrev)
		}

		err = SetMigrationVersion(ctx, q, version)
		if err != nil {
			return errors.Wrapf(err, "could not set migration number for migration %d ('%s')", version, abbrev)
		}

		return nil
	})
}

// SetMigrationVersion records the given version number as the current version
func SetMigrationVersion(ctx context.Context, db Queryable, version int) error {
	_, err := db.Exec(ctx, `
		INSERT INTO `+MIGRATION_TABLE+` VALUES ($1)
	`, version)
	return err
}
