package http_transport

import (
	"account/internal/application/handler/account"
	"account/internal/application/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jmwri/go-http"
)

func Bootstrap(r *chi.Mux, accountService service.Account) {
	getAccountsHandler := account.NewGetAccountsHandler(accountService)
	getAccountHandler := account.NewGetAccountHandler(accountService)
	createAccountHandler := account.NewCreateAccountHandler(accountService)
	updateAccountHandler := account.NewUpdateAccountHandler(accountService)

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(2))
	r.Use(middleware.RequestID)
	r.Use(middleware.RedirectSlashes)

	r.Get("/", go_http.WrapEndpoint(
		getAccountsHandler,
		GetAccountsDecoder,
		GetAccountsEncoder,
		GetAccountsResponder,
	))
	r.Post("/", go_http.WrapEndpoint(
		createAccountHandler,
		CreateAccountDecoder,
		CreateAccountEncoder,
		CreateAccountResponder,
	))
	r.Get("/{id}", go_http.WrapEndpoint(
		getAccountHandler,
		GetAccountDecoder,
		GetAccountEncoder,
		GetAccountResponder,
	))
	r.Put("/{id}", go_http.WrapEndpoint(
		updateAccountHandler,
		UpdateAccountDecoder,
		UpdateAccountEncoder,
		UpdateAccountResponder,
	))
}
