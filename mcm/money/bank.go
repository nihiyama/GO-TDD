package money

// Bank is ...
type Bank struct {
}

// NewBank is ...
func NewBank() *Bank {
	return &Bank{}
}

func (bank *Bank) reduce(expression Expression, to string) *Money {
	return NewDoller(10)
}
