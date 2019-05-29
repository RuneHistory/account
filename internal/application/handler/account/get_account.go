package account

import (
	"account/internal/application/service"
	"account/internal/domain/account"
	"account/internal/mapper"
	"account/internal/transport/http_transport"
	"github.com/go-chi/chi"
	"net/http"
)

type GetAccountRequest struct {
	ID string
}

type GetAccountResponse struct {
	Account *account.Account
}

func NewGetAccountHandler(accountService service.Account) *GetAccountHandler {
	return &GetAccountHandler{
		AccountService: accountService,
	}
}

type GetAccountHandler struct {
	AccountService service.Account
}

func (h *GetAccountHandler) HandleHTTP(w http.ResponseWriter, r *http.Request) {
	req := &GetAccountRequest{
		ID: chi.URLParam(r, "id"),
	}
	res, err := h.handle(req)
	if err != nil {
		http_transport.SendError(err, w)
		return
	}
	mapped := mapper.AccountToHttpV1(res.Account)

	http_transport.SendJson(mapped, w)
}

func (h *GetAccountHandler) handle(r *GetAccountRequest) (*GetAccountResponse, error) {
	acc, err := h.AccountService.GetById(r.ID)
	if err != nil {
		return nil, err
	}
	return &GetAccountResponse{
		Account: acc,
	}, nil
}
