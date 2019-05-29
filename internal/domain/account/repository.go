package account

type Repository interface {
	Get() ([]*Account, error)
	GetById(id string) (*Account, error)
}
