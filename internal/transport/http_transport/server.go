package http_transport

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

func NewServer(address string, accountHandler http.Handler) *Server {
	return &Server{
		httpServer: http.Server{
			Addr: address,
		},
		AccountHandler: accountHandler,
	}
}

type Server struct {
	httpServer     http.Server
	AccountHandler http.Handler
}

func (s *Server) Start(stopWg *sync.WaitGroup, shutdownCh chan struct{}, errCh chan error) {
	stopWg.Add(1)
	defer stopWg.Done()

	startFuncErrCh := make(chan error)
	startFunc := func() {
		fmt.Println("Starting server")
		listener, err := net.Listen("tcp", s.httpServer.Addr)
		if err != nil {
			startFuncErrCh <- err
			return
		}
		fmt.Println("Server listening at " + listener.Addr().String())
		err = s.httpServer.Serve(listener)
		if err != nil {
			startFuncErrCh <- err
			return
		}
	}

	stopFunc := func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := s.httpServer.Shutdown(ctx)
		if err != nil {
			panic(err)
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
