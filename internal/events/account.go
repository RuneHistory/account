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

func RenameAccount(account *account.Account) Event {
	return &RenameAccountEvent{
		Account:   account,
		UpdatedAt: time.Now(),
	}
}

type RenameAccountEvent struct {
	Account   *account.Account
	UpdatedAt time.Time
}

type RenameAccountEventBody struct {
	ID        string    `json:"id"`
	Slug      string    `json:"slug"`
	Nickname  string    `json:"nickname"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (e *RenameAccountEvent) Body() interface{} {
	return RenameAccountEventBody{
		ID:        e.Account.ID,
		Slug:      e.Account.Slug,
		Nickname:  e.Account.Nickname,
		UpdatedAt: e.UpdatedAt,
	}
}

func (e *RenameAccountEvent) Topic() string {
	return "queue.account.rename"
}
