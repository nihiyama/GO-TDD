package money

// NewFranc ...
func NewFranc(amount int) *Money {
	currency := "CHF"
	return &Money{amount, currency}
}
