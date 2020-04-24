package money

import "reflect"

// Doller ...
type Doller struct {
	Money
}

// NewDoller ...
func NewDoller(amount int) *Doller {
	currency := "USD"
	return &Doller{Money{amount, currency}}
}

// GetCurrency is ...
func (doller *Doller) GetCurrency() interface{} {
	return reflect.Indirect(reflect.ValueOf(doller)).Interface()
}

// GetCurrencyName is ...
func (doller *Doller) GetCurrencyName() string {
	return "USD"
}
