package account

import (
	"account/internal/application/service"
	"account/internal/domain/account"
	"account/internal/errs"
	"account/internal/mapper"
	"account/internal/transport/http_transport"
	"net/http"
)

type CreateAccountRequest struct {
	Nickname string `json:"nickname"`
}

type CreateAccountResponse struct {
	Account *account.Account
}

func NewCreateAccountHandler(accountService service.Account) *CreateAccountHandler {
	return &CreateAccountHandler{
		AccountService: accountService,
	}
}

type CreateAccountHandler struct {
	AccountService service.Account
}

func (h *CreateAccountHandler) HandleHTTP(w http.ResponseWriter, r *http.Request) {
	req := &CreateAccountRequest{}
	err := http_transport.ParseJsonBody(r, req)
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

func (h *CreateAccountHandler) handle(r *CreateAccountRequest) (*CreateAccountResponse, error) {
	acc, err := h.AccountService.Create(r.Nickname)
	if err != nil {
		return nil, err
	}
	return &CreateAccountResponse{
		Account: acc,
	}, nil
}
