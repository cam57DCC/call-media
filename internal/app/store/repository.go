package store

import (
	"database/sql"
	"github.com/cam57DCC/call-media/internal/app/model"
)

type RequestRepository interface {
	Add(*model.Request) error
	Update(request *model.Request) (*sql.Tx, error)
}

type ResponseHeaderRepository interface {
	Add(*sql.Tx, *model.ResponseHeader) error
	GetCount(string, string) (*model.CountRequests, error)
}
