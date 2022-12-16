package service

import "os"

type ConfigType struct {
	BindAddr    string
	DatabaseURL string
	AMQPURL     string
}

func NewConfig() *ConfigType {
	return &ConfigType{
		BindAddr:    ":8080",
		DatabaseURL: os.Getenv("DATABASE_URL"),
		AMQPURL:     os.Getenv("AMQP_URL"),
	}
}
