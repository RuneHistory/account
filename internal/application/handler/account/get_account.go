package account

import (
	"account/internal/domain/account"
	"account/internal/mapper"
	"account/internal/transport/http_transport"
	"net/http"
)

type GetAccountRequest struct {
	ID string
}

type GetAccountResponse struct {
	Account *account.Account
}

type GetAccountHandler struct {
}

func (h *GetAccountHandler) HandleHTTP(id string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := &GetAccountRequest{
			ID: id,
		}
		res, err := h.handle(req)
		if err != nil {
			http_transport.SendError(err, w)
			return
		}
		mapped := mapper.AccountToHttpV1(res.Account)

		http_transport.SendJson(mapped, w)
	})
}

func (h *GetAccountHandler) handle(r *GetAccountRequest) (*GetAccountResponse, error) {
	acc := account.NewAccount("1-2-3-4", "Test Account 1", "test-account-1")
	return &GetAccountResponse{
		Account: acc,
	}, nil
}
