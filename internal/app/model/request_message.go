package model

type RequestMessage struct {
	Request
	Second bool `json:"second"`
}
