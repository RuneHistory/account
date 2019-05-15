package account

import (
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

type GetAccountsHandler struct {
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
	accounts := []*account.Account{
		account.NewAccount("1-2-3-4", "Test Account 1", "test-account-1"),
		account.NewAccount("5-6-7-8", "Test Account 2", "test-account-2"),
	}
	return &GetAccountsResponse{
		Accounts: accounts,
	}, nil
}
