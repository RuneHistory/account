package service

import "account/internal/domain/account"

type Account interface {
	Get() ([]*account.Account, error)
	GetById(id string) (*account.Account, error)
}

func NewAccountService(repo account.Repository) Account {
	return &AccountService{
		AccountRepo: repo,
	}
}

type AccountService struct {
	AccountRepo account.Repository
}

func (s *AccountService) Get() ([]*account.Account, error) {
	return s.AccountRepo.Get()
}

func (s *AccountService) GetById(id string) (*account.Account, error) {
	return s.AccountRepo.GetById(id)
}
