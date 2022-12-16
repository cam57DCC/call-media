package sqlstore

import (
	"database/sql"
	"github.com/cam57DCC/call-media/internal/app/store"
	//_ "github.com/lib/pq"
)

type Store struct {
	db                       *sql.DB
	requestRepository        *RequestRepository
	responseHeaderRepository *ResponseHeaderRepository
}

var SQLStore = &Store{}

func New(db *sql.DB) {
	SQLStore.db = db
}

func (s *Store) Request() store.RequestRepository {
	if s.requestRepository != nil {
		return s.requestRepository
	}

	s.requestRepository = &RequestRepository{
		store: s,
	}

	return s.requestRepository
}

func (s *Store) ResponseHeader() store.ResponseHeaderRepository {
	if s.responseHeaderRepository != nil {
		return s.responseHeaderRepository
	}

	s.responseHeaderRepository = &ResponseHeaderRepository{
		store: s,
	}

	return s.responseHeaderRepository
}
