package account

import (
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

type GetAccountHandler struct {
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
	acc := account.NewAccount(r.ID, "Test Account 1", "test-account-1")
	return &GetAccountResponse{
		Account: acc,
	}, nil
}
