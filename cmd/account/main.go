package main

import (
	"account/internal/application/handler"
	"account/internal/transport/http_transport"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	wg := &sync.WaitGroup{}
	shutdownCh := make(chan struct{})
	errCh := make(chan error)
	go handleShutdownSignal(shutdownCh)

	s := buildServer()
	handler.InitHTTP(s)
	go s.Start(wg, shutdownCh, errCh)

	select {
	case err := <-errCh:
		log.Printf("Failed to start up: %s\n", err)

	case <-shutdownCh:
		log.Println("Waiting for shutdown")
		wg.Wait()
		log.Println("All services shutdown")
	}
}

func handleShutdownSignal(shutdownCh chan struct{}) {
	quitCh := make(chan os.Signal)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGTERM)

	hit := false
	for {
		<-quitCh
		if hit {
			os.Exit(0)
		}
		if !hit {
			close(shutdownCh)
		}
		hit = true
	}
}

func buildServer() *http_transport.Server {
	accountHandler := &handler.AccountHandler{}
	return http_transport.NewServer("127.0.0.1:8000", accountHandler)
}
