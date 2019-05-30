package account

import (
	"account/internal/application/service"
	"account/internal/domain/account"
	"account/internal/errs"
	"account/internal/mapper"
	"account/internal/transport/http_transport"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func NewUpdateAccountHandler(accountService service.Account) *UpdateAccountHandler {
	return &UpdateAccountHandler{
		AccountService: accountService,
	}
}

type UpdateAccountHandler struct {
	AccountService service.Account
}

type UpdateAccountRequest struct {
	ID       string
	Nickname string `json:"nickname"`
}

type UpdateAccountResponse struct {
	Account *account.Account
}

func (h *UpdateAccountHandler) HandleHTTP(w http.ResponseWriter, r *http.Request) {
	req := &UpdateAccountRequest{
		ID: chi.URLParam(r, "id"),
	}
	err := http_transport.ParseJsonBody(r, req)
	log.Println(req)
	if err != nil {
		http_transport.SendError(errs.BadRequest(err.Error()), w)
	}

	res, err := h.handle(req)
	if err != nil {
		http_transport.SendError(err, w)
		return
	}
	mapped := mapper.AccountToHttpV1(res.Account)

	http_transport.SendJson(mapped, w)
}

func (h *UpdateAccountHandler) handle(r *UpdateAccountRequest) (*UpdateAccountResponse, error) {
	acc, err := h.AccountService.GetById(r.ID)
	if err != nil {
		return nil, err
	}
	if acc == nil {
		return nil, errs.NotFound(fmt.Sprintf("Account %s not found", r.ID))
	}

	acc.Nickname = r.Nickname

	acc, err = h.AccountService.Update(acc)
	if err != nil {
		return nil, err
	}
	return &UpdateAccountResponse{
		Account: acc,
	}, nil
}
