package controller

import (
	"encoding/json"
	"fmt"
	"github.com/cam57DCC/call-media/internal/app/model"
	"github.com/cam57DCC/call-media/internal/app/service"
	"github.com/cam57DCC/call-media/internal/app/store/sqlstore"
	"github.com/streadway/amqp"
	"net/http"
)

func AddRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var request model.Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "error get body"}`))
		return
	}
	if err = sqlstore.SQLStore.Request().Add(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "error save request"}`))
		return
	}

	rm := model.RequestMessage{
		Request: request,
		Second:  false,
	}

	byteMessage, err := json.Marshal(rm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "internal server error"}`))
		return
	}

	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        byteMessage,
	}

	fmt.Printf("%v\n", service.ChannelMessageBroker)
	if err := service.ChannelMessageBroker.AMPQ.Publish(
		"",
		"SendRequest",
		false,
		false,
		message,
	); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "internal server error"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(`{"status": true}`))
}

func RequestRepeat(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var request model.Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "error get body"}`))
		return
	}
	if err = sqlstore.SQLStore.Request().Add(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "error save request"}`))
		return
	}

	rm := model.RequestMessage{
		Request: request,
		Second:  true,
	}

	byteMessage, err := json.Marshal(rm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "internal server error"}`))
		return
	}

	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        byteMessage,
	}

	fmt.Printf("%v\n", service.ChannelMessageBroker)
	if err := service.ChannelMessageBroker.AMPQ.Publish(
		"",
		"SendRequest",
		false,
		false,
		message,
	); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "internal server error"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(`{"status": true}`))
}

func GetCountRequests(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	header := keys.Get("header")
	headerValue := keys.Get("header_value")
	count, err := sqlstore.SQLStore.ResponseHeader().GetCount(header, headerValue)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "internal server error"}`))
		return
	}
	response, err := json.Marshal(count)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "internal server error"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}
