package events

import (
	"account/internal/domain/account"
	"time"
)

func NewAccount(account *account.Account) Event {
	return &NewAccountEvent{
		Account: account,
	}
}

type NewAccountEvent struct {
	Account *account.Account
}

type NewAccountEventBody struct {
	ID        string    `json:"id"`
	Slug      string    `json:"slug"`
	Nickname  string    `json:"nickname"`
	CreatedAt time.Time `json:"created_at"`
}

func (e *NewAccountEvent) Body() interface{} {
	return NewAccountEventBody{
		ID:        e.Account.ID,
		Slug:      e.Account.Slug,
		Nickname:  e.Account.Nickname,
		CreatedAt: e.Account.CreatedAt,
	}
}

func (e *NewAccountEvent) Topic() string {
	return "queue.account.new"
}
