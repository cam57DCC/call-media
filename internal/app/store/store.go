package store

type Store interface {
	Request() RequestRepository
	Header() ResponseHeaderRepository
}
