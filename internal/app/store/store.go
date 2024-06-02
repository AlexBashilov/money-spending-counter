package store

type Store interface {
	Booker() BookerRepository
}
