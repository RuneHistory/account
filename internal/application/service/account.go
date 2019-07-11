package service

import (
	"account/internal/domain/account"
	"account/internal/domain/validate"
	"account/internal/errs"
	"github.com/RuneHistory/events"
	saramaEvents "github.com/jmwri/go-events/sarama"
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

func NewAccountService(repo account.Repository, validator validate.AccountValidator, publisher *saramaEvents.Publisher) Account {
	return &AccountService{
		AccountRepo: repo,
		Validator:   validator,
		Publisher:   publisher,
	}
}

type AccountService struct {
	AccountRepo account.Repository
	Validator   validate.AccountValidator
	Publisher   *saramaEvents.Publisher
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

	newAccountEvent := &events.NewAccountEvent{
		ID:        acc.ID,
		Slug:      acc.Slug,
		Nickname:  acc.Nickname,
		CreatedAt: acc.CreatedAt,
	}
	err = s.Publisher.Write(newAccountEvent)
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

	renameAccountEvent := &events.RenameAccountEvent{
		ID:        acc.ID,
		Slug:      acc.Slug,
		Nickname:  acc.Nickname,
		UpdatedAt: time.Now(),
	}
	err = s.Publisher.Write(renameAccountEvent)
	if err != nil {
		return nil, errs.InternalServer(err.Error())
	}
	return acc, nil
}
