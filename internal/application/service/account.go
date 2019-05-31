package service

import (
	"account/internal/domain/account"
	"account/internal/domain/validate"
	"account/internal/errs"
	"github.com/mozillazg/go-slugify"
	"github.com/satori/go.uuid"
)

type Account interface {
	Get() ([]*account.Account, error)
	GetById(id string) (*account.Account, error)
	Create(nickname string) (*account.Account, error)
	Update(account *account.Account) (*account.Account, error)
}

func NewAccountService(repo account.Repository, validator validate.AccountValidator) Account {
	return &AccountService{
		AccountRepo: repo,
		Validator:   validator,
	}
}

type AccountService struct {
	AccountRepo account.Repository
	Validator   validate.AccountValidator
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
	if err := s.Validator.NewAccount(a); err != nil {
		return nil, errs.BadRequest(err.Error())
	}
	return s.AccountRepo.Create(a)
}

func (s *AccountService) Update(a *account.Account) (*account.Account, error) {
	a.Slug = slugify.Slugify(a.Nickname)
	if err := s.Validator.UpdateAccount(a); err != nil {
		return nil, errs.BadRequest(err.Error())
	}
	return s.AccountRepo.Update(a)
}
