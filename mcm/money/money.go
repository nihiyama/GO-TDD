package money

import (
	"reflect"
)

// Money is ...
type Money struct {
	amount int
}

// Currency is ...
type Currency interface {
	GetCurrency() interface{}
}

// GetCurrency is ...
func (money *Money) GetCurrency() interface{} {
	return reflect.Indirect(reflect.ValueOf(money)).Interface()
}

// Equals is ..
func Equals(currency Currency, anotherCurrency Currency) bool {
	rfValue := reflect.Indirect(reflect.ValueOf(currency.GetCurrency()))
	rfAnotherValue := reflect.Indirect(reflect.ValueOf(anotherCurrency.GetCurrency()))
	rfAmount := rfValue.FieldByName("Money").Interface().(Money).amount
	rfAnotherAmount := rfAnotherValue.FieldByName("Money").Interface().(Money).amount
	return rfAmount == rfAnotherAmount && rfValue.Type() == rfAnotherValue.Type()
}
