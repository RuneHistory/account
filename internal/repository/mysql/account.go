package mysql

import (
	"account/internal/domain/account"
	"database/sql"
)

func NewAccountMySQL(db *sql.DB) *AccountMySQL {
	return &AccountMySQL{
		DB: db,
	}
}

type AccountMySQL struct {
	DB *sql.DB
}

func (r *AccountMySQL) Get() ([]*account.Account, error) {
	var accounts []*account.Account
	results, err := r.DB.Query("SELECT id, nickname, slug FROM accounts")
	if err == sql.ErrNoRows {
		return accounts, nil
	}
	if err != nil {
		return accounts, err
	}
	for results.Next() {
		var a account.Account
		err = results.Scan(&a.ID, &a.Nickname, &a.Slug)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, &a)
	}
	return accounts, nil
}

func (r *AccountMySQL) GetById(id string) (*account.Account, error) {
	var a account.Account
	err := r.DB.QueryRow("SELECT id, nickname, slug FROM accounts where id = ?", id).Scan(&a.ID, &a.Nickname, &a.Slug)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &a, nil
}
