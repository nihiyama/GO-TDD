package money

// Money is ...
type Money struct {
	amount int
}

// Equals is ...
func (money *Money) equals(anotherMoney *Money) bool {
	return money.amount == anotherMoney.amount
}
