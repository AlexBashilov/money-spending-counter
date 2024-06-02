package sqlstore

import (
	"booker/internal/app/store"
	"database/sql"
)

// Store ...
type Store struct {
	db               *sql.DB
	bookerRepository *BookerRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Booker() store.BookerRepository {
	if s.bookerRepository != nil {
		return s.bookerRepository
	}

	s.bookerRepository = &BookerRepository{
		store: s,
	}

	return s.bookerRepository
}
