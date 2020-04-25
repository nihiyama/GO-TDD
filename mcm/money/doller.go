package money

// NewDoller ...
func NewDoller(amount int) *Money {
	currency := "USD"
	return &Money{amount, currency}
}
