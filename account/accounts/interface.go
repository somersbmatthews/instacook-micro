package accounts

type Accounter interface {
	Create(acc *Account) error
	Update(acc *Account) error
}
