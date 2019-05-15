package main

import (
	"account/internal/application/handler/account"
	"account/internal/transport/http_transport"
	"errors"
	"github.com/go-chi/chi"
	_ "github.com/go-chi/chi"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	address := os.Getenv("LISTEN_ADDRESS")

	wg := &sync.WaitGroup{}
	shutdownCh := make(chan struct{})
	errCh := make(chan error)
	go handleShutdownSignal(errCh)

	r := chi.NewRouter()

	getAccountsHandler := &account.GetAccountsHandler{}
	getAccountHandler := &account.GetAccountHandler{}

	r.Get("/", getAccountsHandler.HandleHTTP)
	r.Get("/{id}", getAccountHandler.HandleHTTP)

	go http_transport.Start(address, r, wg, shutdownCh, errCh)

	err := <-errCh
	if err != nil {
		log.Printf("fatal err: %s\n", err)
	}

	log.Println("initiating graceful shutdown")
	close(shutdownCh)

	wg.Wait()
	log.Println("shutdown")
}

func handleShutdownSignal(errCh chan error) {
	quitCh := make(chan os.Signal)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGTERM)

	hit := false
	for {
		<-quitCh
		if hit {
			os.Exit(0)
		}
		if !hit {
			errCh <- errors.New("shutdown signal received")
		}
		hit = true
	}
}
