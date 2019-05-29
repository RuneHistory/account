package http_transport

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

func Start(address string, r http.Handler, stopWg *sync.WaitGroup, shutdownCh chan struct{}, errCh chan error) {
	stopWg.Add(1)
	defer stopWg.Done()

	s := &http.Server{
		Addr:    address,
		Handler: r,
	}

	startFuncErrCh := make(chan error)
	startFunc := func() {
		log.Println("Starting server")
		listener, err := net.Listen("tcp", s.Addr)
		if err != nil {
			startFuncErrCh <- err
			return
		}
		log.Printf("Server listening at %s\n", listener.Addr().String())
		err = s.Serve(listener)
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Serve failed: %s", err)
		}
	}

	stopFunc := func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := s.Shutdown(ctx)
		if err != nil {
			log.Fatalf("Shutdown failed: %s", err)
		}
	}

	go startFunc()

	select {
	case err := <-startFuncErrCh:
		errCh <- err
	case <-shutdownCh:
		stopFunc()
	}
}
