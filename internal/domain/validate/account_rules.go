package validate

import (
	"account/internal/domain/account"
	"errors"
	"fmt"
)

const (
	IDLength          = 36
	MaxNicknameLength = 12
	MaxSlugLength     = 12
)

type AccountRules interface {
	IDIsPresent(a *account.Account) error
	IDIsCorrectLength(a *account.Account) error
	IDWillBeUnique(a *account.Account) error
	IDIsUnique(a *account.Account) error
	NicknameIsPresent(a *account.Account) error
	NicknameIsNotTooLong(a *account.Account) error
	NicknameIsUniqueToID(a *account.Account) error
	SlugIsPresent(a *account.Account) error
	SlugIsNotTooLong(a *account.Account) error
	SlugIsUniqueToID(a *account.Account) error
}

func NewAccountRules(accountRepo account.Repository) AccountRules {
	return &StdAccountRules{
		AccountRepo: accountRepo,
	}
}

type StdAccountRules struct {
	AccountRepo account.Repository
}

func (x *StdAccountRules) IDIsPresent(a *account.Account) error {
	if a.ID == "" {
		return errors.New("id is blank")
	}
	return nil
}

func (x *StdAccountRules) IDIsCorrectLength(a *account.Account) error {
	if len(a.ID) != IDLength {
		return fmt.Errorf("id %s must be exactly %d characters", a.ID, IDLength)
	}
	return nil
}

func (x *StdAccountRules) IDWillBeUnique(a *account.Account) error {
	amount, err := x.AccountRepo.CountId(a.ID)
	if err != nil {
		return err
	}
	if amount != 0 {
		return fmt.Errorf("ID %s must be unique", a.ID)
	}
	return nil
}

func (x *StdAccountRules) IDIsUnique(a *account.Account) error {
	count, err := x.AccountRepo.CountId(a.ID)
	if err != nil {
		return err
	}
	if count != 1 {
		return fmt.Errorf("ID %s must be unique", a.ID)
	}
	return nil
}

func (x *StdAccountRules) NicknameIsPresent(a *account.Account) error {
	if a.Nickname == "" {
		return errors.New("nickname is blank")
	}
	return nil
}

func (x *StdAccountRules) NicknameIsNotTooLong(a *account.Account) error {
	if len(a.Nickname) > MaxNicknameLength {
		return fmt.Errorf("nickname must be no longer than %d characters", MaxNicknameLength)
	}
	return nil
}

func (x *StdAccountRules) NicknameIsUniqueToID(a *account.Account) error {
	acc, err := x.AccountRepo.GetByNicknameWithoutId(a.Nickname, a.ID)
	if err != nil {
		return err
	}
	if acc != nil {
		return fmt.Errorf("nickname %s already exists", a.Nickname)
	}
	return nil
}

func (x *StdAccountRules) SlugIsPresent(a *account.Account) error {
	if a.Slug == "" {
		return errors.New("slug is blank")
	}
	return nil
}

func (x *StdAccountRules) SlugIsNotTooLong(a *account.Account) error {
	if len(a.Slug) > MaxSlugLength {
		return fmt.Errorf("slug must be no longer than %d characters", MaxSlugLength)
	}
	return nil
}

func (x *StdAccountRules) SlugIsUniqueToID(a *account.Account) error {
	acc, err := x.AccountRepo.GetBySlugWithoutId(a.Slug, a.ID)
	if err != nil {
		return err
	}
	if acc != nil {
		return fmt.Errorf("slug %s already exists", a.Slug)
	}
	return nil
}
