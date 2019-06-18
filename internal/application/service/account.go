package service

import (
	"account/internal/domain/account"
	"account/internal/domain/validate"
	"account/internal/errs"
	"account/internal/events"
	"github.com/mozillazg/go-slugify"
	"github.com/satori/go.uuid"
	"time"
)

type Account interface {
	Get() ([]*account.Account, error)
	GetById(id string) (*account.Account, error)
	Create(nickname string) (*account.Account, error)
	Update(account *account.Account) (*account.Account, error)
}

func NewAccountService(repo account.Repository, validator validate.AccountValidator, dispatcher events.Dispatcher) Account {
	return &AccountService{
		AccountRepo: repo,
		Validator:   validator,
		Dispatcher:  dispatcher,
	}
}

type AccountService struct {
	AccountRepo account.Repository
	Validator   validate.AccountValidator
	Dispatcher  events.Dispatcher
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
	now := time.Now()
	a := account.NewAccount(id, nickname, slug, now)
	if err := s.Validator.NewAccount(a); err != nil {
		return nil, errs.BadRequest(err.Error())
	}
	acc, err := s.AccountRepo.Create(a)
	if err != nil {
		return nil, errs.InternalServer(err.Error())
	}

	event := events.NewAccount(acc)
	err = s.Dispatcher.Dispatch(event)
	if err != nil {
		return nil, errs.InternalServer(err.Error())
	}
	return acc, nil
}

func (s *AccountService) Update(a *account.Account) (*account.Account, error) {
	a.Slug = slugify.Slugify(a.Nickname)
	if err := s.Validator.UpdateAccount(a); err != nil {
		return nil, errs.BadRequest(err.Error())
	}
	acc, err := s.AccountRepo.Update(a)
	if err != nil {
		return nil, errs.InternalServer(err.Error())
	}

	event := events.RenameAccount(acc)
	err = s.Dispatcher.Dispatch(event)
	if err != nil {
		return nil, errs.InternalServer(err.Error())
	}
	return acc, nil
}
