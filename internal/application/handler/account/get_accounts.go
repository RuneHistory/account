package account

import (
	"account/internal/application/service"
	"account/internal/domain/account"
)

type GetAccountsRequest struct {
}

type GetAccountsResponse struct {
	Accounts []*account.Account
}

func NewGetAccountsHandler(accountService service.Account) *GetAccountsHandler {
	return &GetAccountsHandler{
		AccountService: accountService,
	}
}

type GetAccountsHandler struct {
	AccountService service.Account
}

func (h *GetAccountsHandler) Handle(r interface{}) (interface{}, error) {
	accounts, err := h.AccountService.Get()
	if err != nil {
		return nil, err
	}
	return &GetAccountsResponse{
		Accounts: accounts,
	}, nil
}
