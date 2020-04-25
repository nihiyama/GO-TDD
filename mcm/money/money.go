package money

// Money is ...
type Money struct {
	amount   int
	currency string
}

// // Currency is ...
// type Currency interface {
// 	GetCurrency() interface{}
// }

// NewMoney is ...
func NewMoney(amount int, currency ...string) *Money {
	if len(currency) == 0 {
		currency = append(currency, "None")
	}
	return &Money{amount, currency[0]}
}

// GetCurrency is ...
// func (money *Money) GetCurrency() interface{} {
// 	return reflect.Indirect(reflect.ValueOf(money)).Interface()
// }

// Times is ..
func (money *Money) times(multiple int) *Money {
	return NewMoney(money.amount * multiple)
}

func (money *Money) plus(add *Money) *Money {
	return NewMoney(money.amount+add.amount, money.currency)
}

func (money *Money) equals(anotheMoney *Money) bool {
	return money.amount == anotheMoney.amount &&
		money.currency == anotheMoney.currency
}

// Equals is ..
// func Equals(currency Currency, anotherCurrency Currency) bool {
// 	rfValue := reflect.Indirect(reflect.ValueOf(currency.GetCurrency()))
// 	rfAnotherValue := reflect.Indirect(reflect.ValueOf(anotherCurrency.GetCurrency()))
// 	rfMoney := rfValue.Interface().(Money)
// 	rfAnotherMoney := rfAnotherValue.Interface().(Money)
// 	return rfMoney.amount == rfAnotherMoney.amount && rfMoney.currency == rfAnotherMoney.currency
// }
