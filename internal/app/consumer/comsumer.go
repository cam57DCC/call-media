package consumer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cam57DCC/call-media/internal/app/apiserver"
	"github.com/cam57DCC/call-media/internal/app/model"
	"github.com/cam57DCC/call-media/internal/app/service"
	"github.com/cam57DCC/call-media/internal/app/store/sqlstore"
	"net/http"
	"os"
	"strings"
	"time"
)

func Run(config *service.ConfigType) error {
	db, err := apiserver.NewDB(config)
	if err != nil {
		return err
	}
	defer db.Close()
	sqlstore.New(db)

	connectRabbitMQ, err := service.NewConsumerChanel(config)
	defer connectRabbitMQ.Close()

	messages, err := service.ChannelMessageBroker.AMPQ.Consume(
		"SendRequest",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	fmt.Println("Successfully connected to RabbitMQ")
	fmt.Println("Waiting for messages")

	forever := make(chan bool)

	go func() {
		for message := range messages {
			time.Sleep(1 * time.Second)
			// For example, show received message in a console.
			fmt.Printf(" > Received message: %s\n", message.Body)
			requestMessage := &model.RequestMessage{}
			if err := json.Unmarshal(message.Body, requestMessage); err != nil {
				fmt.Println(err.Error())
				continue
			}
			client := http.Client{}
			request, err := http.NewRequest("GET", requestMessage.URL, nil)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			resp, err := client.Do(request)
			if err != nil {
				if requestMessage.Second == false {
					time.Sleep(15 * time.Second)
					requestMessage.Second = true
					body, err := json.Marshal(requestMessage)
					if err != nil {
						fmt.Println(err.Error())
						continue
					}
					request, err := http.NewRequest("POST", os.Getenv("REPEAT_REQUEST_URL"), bytes.NewBuffer(body))
					if err != nil {
						fmt.Println(err.Error())
						continue
					}

					_, err = client.Do(request)
					if err != nil {
						fmt.Println(err.Error())
						continue
					}
				}
				fmt.Println(err.Error())
				continue
			}
			var body []byte
			if _, err = resp.Body.Read(body); err != nil {
				fmt.Println(err.Error())
				continue
			}

			var respBody []byte
			_, err = resp.Body.Read(respBody)
			if err != nil {
				fmt.Println(err)
				continue
			}
			requestModel := &model.Request{
				ID:           requestMessage.ID,
				Response:     string(respBody),
				ResponseCode: resp.StatusCode,
			}
			tx, err := sqlstore.SQLStore.Request().Update(requestModel)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			for key, header := range resp.Header {
				responseHeader := &model.ResponseHeader{
					RequestId:   requestModel.ID,
					Header:      key,
					HeaderValue: strings.Join(header, ", "),
				}
				if err = sqlstore.SQLStore.ResponseHeader().Add(tx, responseHeader); err != nil {
					fmt.Println(err.Error())
					break
				}
			}
			if err = tx.Commit(); err != nil {
				fmt.Println(err)
			}
		}
	}()

	<-forever

	return nil
}
