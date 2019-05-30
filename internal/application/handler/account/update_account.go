package account

import (
	"account/internal/application/service"
	"account/internal/domain/account"
	"account/internal/errs"
	"fmt"
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

func (h *UpdateAccountHandler) Handle(r interface{}) (interface{}, error) {
	req := r.(*UpdateAccountRequest)
	acc, err := h.AccountService.GetById(req.ID)
	if err != nil {
		return nil, err
	}
	if acc == nil {
		return nil, errs.NotFound(fmt.Sprintf("Account %s not found", req.ID))
	}

	acc.Nickname = req.Nickname

	acc, err = h.AccountService.Update(acc)
	if err != nil {
		return nil, err
	}
	return &UpdateAccountResponse{
		Account: acc,
	}, nil
}
