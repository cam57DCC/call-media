package model

type CountRequests struct {
	Count         int64 `json:"count"`
	CountByHeader int64 `json:"count_by_header"`
}
