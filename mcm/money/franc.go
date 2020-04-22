package money

import "reflect"

// Franc ...
type Franc struct {
	Money
}

// NewFranc ...
func NewFranc(amount int) *Franc {
	currency := "CHF"
	return &Franc{Money{amount, currency}}
}

// func (franc *Franc) times(multiple int) *Franc {
// 	return &Franc{Money{franc.amount * multiple}}
// }

// GetCurrency is ...
func (franc *Franc) GetCurrency() interface{} {
	return reflect.Indirect(reflect.ValueOf(franc)).Interface()
}

// GetCurrencyName is ...
func (franc *Franc) GetCurrencyName() string {
	return franc.currency
}
