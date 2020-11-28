package db

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

func InTransaction(ctx context.Context, log *zerolog.Logger, db Transactable, f func(Queryable) error) error {
	log.Trace().Msg("beginning transaction")
	tx, txErr := db.Begin(ctx)
	if txErr != nil {
		return errors.Wrap(txErr, "could not begin transaction")
	}

	err := f(tx)

	if err == nil {
		log.Trace().Msg("committing transaction")
		txErr = tx.Commit(ctx)
		if txErr != nil {
			return errors.Wrap(txErr, "could not commit transaction")
		}
	} else {
		log.Trace().Msg("rolling back transaction")
		txErr = tx.Rollback(ctx)
		if txErr != nil {
			return errors.Wrapf(txErr, "could not rollback transaction while handling error '%s'", err.Error())
		}
	}

	return err
}
