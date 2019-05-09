package main

import (
	"account/internal/application/handler"
	"account/internal/transport/http_transport"
	"log"
	"net/http"
)

func main() {
	accountHandler := &handler.AccountHandler{}
	s := http_transport.NewServer(accountHandler)
	err := http.ListenAndServe("127.0.0.1:8000", s)
	if err != nil {
		log.Printf("Failed to start up: %s", err)
	}
}
