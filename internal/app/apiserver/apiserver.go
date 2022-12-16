package apiserver

import (
	"context"
	"database/sql"
	"github.com/cam57DCC/call-media/internal/app/service"
	"github.com/cam57DCC/call-media/internal/app/store/sqlstore"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Start(config *service.ConfigType) error {
	db, err := NewDB(config)
	if err != nil {
		return err
	}
	defer db.Close()
	sqlstore.New(db)

	connMessageBroker, err := service.NewProducer(config)
	if err != nil {
		return err
	}
	defer connMessageBroker.Close()
	defer service.ChannelMessageBroker.AMPQ.Close()

	srv := newServer()

	s := &http.Server{
		Handler: http.TimeoutHandler(srv.Router, 60*time.Second, `{"error": "Превышено время ожидания"}`),
		Addr:    config.BindAddr,
	}

	go func() {
		if err = s.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	wait := time.Second * 16

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	s.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)

	return err
}

func NewDB(config *service.ConfigType) (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DatabaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
