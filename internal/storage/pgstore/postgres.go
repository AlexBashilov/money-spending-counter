package pgstore

import (
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
)

type PGStorage struct {
	dbRO bun.IDB
	dbRW bun.IDB
}

type PgPublisherStorage struct {
	*PGStorage
}

func NewStorage(dbRO, dbRW bun.IDB) *PGStorage {
	return &PGStorage{
		dbRO: dbRO,
		dbRW: dbRW,
	}
}

func NewPGPublisherStorage(repo *PGStorage) *PgPublisherStorage {
	return &PgPublisherStorage{
		PGStorage: repo,
	}
}

func finishTx(err error, tx bun.Tx) error {
	if err != nil {
		_ = tx.Rollback()

		return err
	}

	if commitErr := tx.Commit(); commitErr != nil {
		return errors.Wrap(err, "finishTx")
	}

	return nil
}
