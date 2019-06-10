package account

import (
	"account/internal/application/service"
	"account/internal/domain/account"
	"account/internal/errs"
	"fmt"
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
	if acc == nil {
		return nil, errs.NotFound(fmt.Sprintf("Account %s not found", req.ID))
	}
	return &GetAccountResponse{
		Account: acc,
	}, nil
}
