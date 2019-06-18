package account

type Publisher interface {
	New(a *Account) error
}
