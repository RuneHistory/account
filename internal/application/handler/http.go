package handler

import (
	"account/internal/errs"
	"account/internal/transport/http_transport"
	"fmt"
	"net/http"
)

func InitHTTP(server *http_transport.Server) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var head string
		head, r.URL.Path = http_transport.ShiftPath(r.URL.Path)

		switch head {
		case "accounts":
			server.AccountHandler.ServeHTTP(w, r)
			return
		default:
			http_transport.SendError(errs.NotFound(fmt.Sprintf("Unknown resource: %s", head)), w)
		}
	})
}
