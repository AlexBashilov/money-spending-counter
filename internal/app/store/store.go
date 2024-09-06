package store

// Store interface
type Store interface {
	Booker() BookerRepository
}
