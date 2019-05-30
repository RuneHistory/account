package account

import (
	"account/internal/application/service"
	"account/internal/domain/account"
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

func (h *GetAccountHandler) Handle(r interface{}) (interface{}, error) {
	req := r.(*GetAccountRequest)
	acc, err := h.AccountService.GetById(req.ID)
	if err != nil {
		return nil, err
	}
	return &GetAccountResponse{
		Account: acc,
	}, nil
}
