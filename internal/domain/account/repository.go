package account

type Repository interface {
	Get() ([]*Account, error)
	GetById(id string) (*Account, error)
	Create(a *Account) (*Account, error)
	Update(a *Account) (*Account, error)
}
