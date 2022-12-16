package main

import (
	"github.com/cam57DCC/call-media/internal/app/consumer"
	"github.com/cam57DCC/call-media/internal/app/service"
	"log"
)

func main() {
	config := service.NewConfig()
	if err := consumer.Run(config); err != nil {
		log.Fatalln(err)
	}
}
