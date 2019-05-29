package service

import (
	"account/internal/domain/account"
	"github.com/mozillazg/go-slugify"
	"github.com/satori/go.uuid"
)

type Account interface {
	Get() ([]*account.Account, error)
	GetById(id string) (*account.Account, error)
	Create(nickname string) (*account.Account, error)
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

func (s *AccountService) Create(nickname string) (*account.Account, error) {
	id := uuid.NewV4().String()
	slug := slugify.Slugify(nickname)
	a := account.NewAccount(id, nickname, slug)
	// TODO: Implement validation
	return s.AccountRepo.Create(a)
}
