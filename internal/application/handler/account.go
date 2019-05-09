package handler

import (
	"account/internal/application/handler/account"
	"account/internal/errs"
	"account/internal/transport/http_transport"
	"net/http"
)

type AccountHandler struct {
	getAccountsHandler account.GetAccountsHandler
	getAccountHandler account.GetAccountHandler
}

func (h *AccountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = http_transport.ShiftPath(r.URL.Path)
	if head == "" && r.Method == "GET" {
		h.getAccountsHandler.HandleHTTP().ServeHTTP(w, r)
		return
	}
	switch r.Method {
	case "GET":
		h.getAccountHandler.HandleHTTP(head).ServeHTTP(w, r)
	default:
		http_transport.SendError(errs.MethodNotAllowed("Only GET is allowed"), w)
	}
}
