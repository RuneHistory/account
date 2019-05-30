package http_transport

import (
	"account/internal/application/handler/account"
	"account/internal/application/service"
	"account/internal/errs"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func WrapEndpoint(endpoint Endpoint, decoder DecoderFunc, encoder EncoderFunc, responder ResponderFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, err := decoder(r)
		if err != nil {
			SendError(errs.BadRequest(err.Error()), w)
			return
		}
		res, err := endpoint.Handle(req)
		if err != nil {
			SendError(err, w)
			return
		}
		encoded, err := encoder(res)
		if err != nil {
			SendError(err, w)
			return
		}
		responder(w, encoded)
	})
}

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

	r.Get("/", WrapEndpoint(
		getAccountsHandler,
		GetAccountsDecoder,
		GetAccountsEncoder,
		GetAccountsResponder,
	))
	r.Post("/", WrapEndpoint(
		createAccountHandler,
		CreateAccountDecoder,
		CreateAccountEncoder,
		CreateAccountResponder,
	))
	r.Get("/{id}", WrapEndpoint(
		getAccountHandler,
		GetAccountDecoder,
		GetAccountEncoder,
		GetAccountResponder,
	))
	r.Put("/{id}", WrapEndpoint(
		updateAccountHandler,
		UpdateAccountDecoder,
		UpdateAccountEncoder,
		UpdateAccountResponder,
	))
}
