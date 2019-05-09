package http_transport

import (
	"account/internal/errs"
	"fmt"
	"net/http"
)

func NewServer(accountHandler http.Handler) *Server {
	return &Server{
		AccountHandler: accountHandler,
	}
}

type Server struct {
	AccountHandler http.Handler
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = ShiftPath(r.URL.Path)

	switch head {
	case "accounts":
		s.AccountHandler.ServeHTTP(w, r)
		return
	default:
		SendError(errs.NotFound(fmt.Sprintf("Unknown resource: %s", head)), w)
	}
}