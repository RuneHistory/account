package mapper

import (
	"account/internal/domain/account"
	"time"
)

type AccountHttpV1 struct {
	ID        string    `json:"id"`
	Slug      string    `json:"slug"`
	Nickname  string    `json:"nickname"`
	CreatedAt time.Time `json:"created_at"`
}

func AccountToHttpV1(acc *account.Account) *AccountHttpV1 {
	return &AccountHttpV1{
		ID:        acc.ID,
		Slug:      acc.Slug,
		Nickname:  acc.Nickname,
		CreatedAt: acc.CreatedAt,
	}
}

func AccountFromHttpV1(acc *AccountHttpV1) *account.Account {
	return &account.Account{
		ID:        acc.ID,
		Slug:      acc.Slug,
		Nickname:  acc.Nickname,
		CreatedAt: acc.CreatedAt,
	}
}
