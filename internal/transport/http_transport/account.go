package http_transport

import (
	accountHandler "account/internal/application/handler/account"
	"account/internal/errs"
	"account/internal/mapper"
	"github.com/go-chi/chi"
	"github.com/jmwri/go-http"
	"net/http"
)

func CreateAccountDecoder(r *http.Request) (interface{}, error) {
	req := &accountHandler.CreateAccountRequest{}
	err := go_http.ParseJsonBody(r, req)
	if err != nil {
		return nil, errs.BadRequest(err.Error())
	}
	return req, nil
}

func CreateAccountEncoder(d interface{}) (interface{}, error) {
	res := d.(*accountHandler.CreateAccountResponse)
	return mapper.AccountToHttpV1(res.Account), nil
}

func CreateAccountResponder(w http.ResponseWriter, d interface{}) {
	go_http.SendJson(d, w)
}

func GetAccountDecoder(r *http.Request) (interface{}, error) {
	req := &accountHandler.GetAccountRequest{
		ID: chi.URLParam(r, "id"),
	}
	return req, nil
}

func GetAccountEncoder(d interface{}) (interface{}, error) {
	res := d.(*accountHandler.GetAccountResponse)
	return mapper.AccountToHttpV1(res.Account), nil
}

func GetAccountResponder(w http.ResponseWriter, d interface{}) {
	go_http.SendJson(d, w)
}

func GetAccountsDecoder(_ *http.Request) (interface{}, error) {
	req := &accountHandler.GetAccountsRequest{}
	return req, nil
}

func GetAccountsEncoder(d interface{}) (interface{}, error) {
	res := d.(*accountHandler.GetAccountsResponse)
	mapped := make([]*mapper.AccountHttpV1, len(res.Accounts))
	for k, acc := range res.Accounts {
		mapped[k] = mapper.AccountToHttpV1(acc)
	}
	return mapped, nil
}

func GetAccountsResponder(w http.ResponseWriter, d interface{}) {
	go_http.SendJson(d, w)
}

func UpdateAccountDecoder(r *http.Request) (interface{}, error) {
	req := &accountHandler.UpdateAccountRequest{
		ID: chi.URLParam(r, "id"),
	}
	err := go_http.ParseJsonBody(r, req)
	if err != nil {
		return nil, errs.BadRequest(err.Error())
	}
	return req, nil
}

func UpdateAccountEncoder(d interface{}) (interface{}, error) {
	res := d.(*accountHandler.UpdateAccountResponse)
	return mapper.AccountToHttpV1(res.Account), nil
}

func UpdateAccountResponder(w http.ResponseWriter, d interface{}) {
	go_http.SendJson(d, w)
}
