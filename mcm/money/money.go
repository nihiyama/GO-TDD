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
func NewMoney(amount int, currency ...string) *Money {
	if len(currency) == 0 {
		currency = append(currency, "None")
	}
	return &Money{amount, currency[0]}
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
	var rfMoney, rfAnotherMoney Money
	if rfValue.Type().Field(0).Name == "Money" {
		rfMoney = rfValue.FieldByName("Money").Interface().(Money)
	} else {
		rfMoney = rfValue.Interface().(Money)
	}
	if rfAnotherValue.Type().Field(0).Name == "Money" {
		rfAnotherMoney = rfAnotherValue.FieldByName("Money").Interface().(Money)
	} else {
		rfAnotherMoney = rfAnotherValue.Interface().(Money)
	}
	return rfMoney.amount == rfAnotherMoney.amount && rfMoney.currency == rfAnotherMoney.currency
}
