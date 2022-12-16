package sqlstore

import (
	"database/sql"
	"github.com/cam57DCC/call-media/internal/app/model"
)

type ResponseHeaderRepository struct {
	store *Store
}

func (r *ResponseHeaderRepository) Add(tx *sql.Tx, header *model.ResponseHeader) error {
	res, err := tx.Exec(
		"INSERT INTO response_headers(request_id, header, header_value) VALUES (?, ?, ?)",
		header.RequestId,
		header.Header,
		header.HeaderValue,
	)
	if err != nil {
		tx.Rollback()
		return err
	}
	header.ID, err = res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (r *ResponseHeaderRepository) GetCount(header, headerValue string) (*model.CountRequests, error) {
	count := &model.CountRequests{}
	row := r.store.db.QueryRow(
		`SELECT (
					SELECT count FROM count_requests LIMIT 1
				), (
					SELECT count(id)
				    FROM response_headers
				    WHERE header = ? AND header_value = ?
				)`,
		header,
		headerValue,
	)
	if err := row.Scan(&count.Count, &count.CountByHeader); err != nil {
		return nil, err
	}
	return count, nil
}
