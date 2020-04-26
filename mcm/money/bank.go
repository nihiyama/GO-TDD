package money

// Bank is ...
type Bank struct {
}

// NewBank is ...
func NewBank() *Bank {
	return &Bank{}
}

func (bank *Bank) reduce(expression Expression, to string) *Money {
	resultAmount := 0
	for _, money := range expression.GetMoney() {
		resultAmount += money.amount
	}
	return NewMoney(resultAmount, to)
}
