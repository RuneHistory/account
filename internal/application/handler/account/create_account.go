package account

import (
	"account/internal/application/service"
	"account/internal/domain/account"
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

func (h *CreateAccountHandler) Handle(r interface{}) (interface{}, error) {
	req := r.(*CreateAccountRequest)
	acc, err := h.AccountService.Create(req.Nickname)
	if err != nil {
		return nil, err
	}
	return &CreateAccountResponse{
		Account: acc,
	}, nil
}
