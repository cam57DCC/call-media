package model

type Request struct {
	ID           int64  `json:"id"`
	URL          string `json:"url"`
	ResponseCode int    `json:"response_code"`
	Response     string `json:"response"`
}
