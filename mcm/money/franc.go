package money

import "reflect"

// Franc ...
type Franc struct {
	Money
}

// NewFranc ...
func NewFranc(amount int) *Franc {
	return &Franc{Money{amount}}
}

func (franc *Franc) times(multiple int) *Franc {
	return &Franc{Money{franc.amount * multiple}}
}

// GetCurrency is ...
func (franc *Franc) GetCurrency() interface{} {
	return reflect.Indirect(reflect.ValueOf(franc)).Interface()
}
