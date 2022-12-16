package sqlstore

import (
	"database/sql"
	"github.com/cam57DCC/call-media/internal/app/model"
)

type RequestRepository struct {
	store *Store
}

func (r *RequestRepository) Add(request *model.Request) error {

	tx, err := r.store.db.Begin()
	if err != nil {
		return err
	}

	res, err := tx.Exec("INSERT INTO requests(url) VALUES (?)", request.URL)
	if err != nil {
		tx.Rollback()
		return err
	}
	request.ID, err = res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	res, err = tx.Exec("UPDATE count_requests SET count = count +1;")
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *RequestRepository) Update(request *model.Request) (*sql.Tx, error) {
	tx, err := r.store.db.Begin()
	if err != nil {
		return nil, err
	}

	res, err := tx.Exec(
		"UPDATE requests SET response_code = ?, response = ? WHERE id = ?",
		request.ResponseCode,
		request.Response,
		request.ID,
	)
	if err != nil {
		return nil, err
	}
	_, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return tx, nil
}
