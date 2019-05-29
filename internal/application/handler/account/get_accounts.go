package account

import (
	"account/internal/application/service"
	"account/internal/domain/account"
	"account/internal/mapper"
	"account/internal/transport/http_transport"
	"net/http"
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

func (h *GetAccountsHandler) HandleHTTP(w http.ResponseWriter, r *http.Request) {
	req := &GetAccountsRequest{}

	res, err := h.handle(req)
	if err != nil {
		http_transport.SendError(err, w)
		return
	}

	mapped := make([]*mapper.AccountHttpV1, len(res.Accounts))
	for k, acc := range res.Accounts {
		mapped[k] = mapper.AccountToHttpV1(acc)
	}

	http_transport.SendJson(mapped, w)
}

func (h *GetAccountsHandler) handle(r *GetAccountsRequest) (*GetAccountsResponse, error) {
	accounts, err := h.AccountService.Get()
	if err != nil {
		return nil, err
	}
	return &GetAccountsResponse{
		Accounts: accounts,
	}, nil
}
