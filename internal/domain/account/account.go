package account

import "time"

type Account struct {
	ID        string
	Nickname  string
	Slug      string
	CreatedAt time.Time
}

func NewAccount(uuid string, nickname string, slug string, createdAt time.Time) *Account {
	return &Account{
		ID:        uuid,
		Nickname:  nickname,
		Slug:      slug,
		CreatedAt: createdAt,
	}
}
