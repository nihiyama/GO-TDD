package money

import "reflect"

// Doller ...
type Doller struct {
	Money
}

// NewDoller ...
func NewDoller(amount int) *Doller {
	return &Doller{Money{amount}}
}

// func (doller *Doller) times(multiple int) *Doller {
// 	return &Doller{Money{doller.amount * multiple}}
// }

// GetCurrency is ...
func (doller *Doller) GetCurrency() interface{} {
	return reflect.Indirect(reflect.ValueOf(doller)).Interface()
}
