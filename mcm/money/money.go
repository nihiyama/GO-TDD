package money

import (
	"reflect"
)

// Money is ...
type Money struct {
	amount   int
	currency string
}

// Currency is ...
type Currency interface {
	GetCurrency() interface{}
	GetCurrencyName() string
	Times(int) *Money
}

// NewMoney is ...
func NewMoney(amount int) *Money {
	currency := "None"
	return &Money{amount, currency}
}

// GetCurrency is ...
func (money *Money) GetCurrency() interface{} {
	return reflect.Indirect(reflect.ValueOf(money)).Interface()
}

// GetCurrencyName is ...
func (money *Money) GetCurrencyName() string {
	return money.currency
}

// Times is ..
func (money *Money) Times(multiple int) *Money {
	return NewMoney(money.amount * multiple)
}

// Equals is ..
func Equals(currency Currency, anotherCurrency Currency) bool {
	rfValue := reflect.Indirect(reflect.ValueOf(currency.GetCurrency()))
	rfAnotherValue := reflect.Indirect(reflect.ValueOf(anotherCurrency.GetCurrency()))
	if rfValue.Type().Field(0).Name == "Money" {
		rfAmount := rfValue.FieldByName("Money").Interface().(Money).amount
		rfAnotherAmount := rfAnotherValue.FieldByName("Money").Interface().(Money).amount
		return rfAmount == rfAnotherAmount && rfValue.Type() == rfAnotherValue.Type()
	}
	rfAmount := rfValue.Interface().(Money).amount
	rfAnotherAmount := rfAnotherValue.Interface().(Money).amount
	return rfAmount == rfAnotherAmount && rfValue.Type() == rfAnotherValue.Type()
}
