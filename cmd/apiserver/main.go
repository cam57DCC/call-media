package main

import (
	"github.com/cam57DCC/call-media/internal/app/apiserver"
	"github.com/cam57DCC/call-media/internal/app/service"
	"log"
)

func main() {
	config := service.NewConfig()
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
