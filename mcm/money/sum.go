package money

import "fmt"

// Sum is ...
type Sum struct {
	augend Money
	addend Money
}

// NewSum is ...
func NewSum(augend Money, addend Money) *Sum {
	return &Sum{augend, addend}
}

// ToString ...
func (sum *Sum) ToString() string {
	return fmt.Sprintf("[%v, %v]", sum.augend, sum.addend)
}

// GetMoney is ...
func (sum *Sum) GetMoney() []Money {
	return []Money{sum.augend, sum.addend}
}

// Reduce is ...
func (sum *Sum) Reduce(bank *Bank, to string) *Money {
	resultAmount := sum.augend.amount + sum.addend.amount
	return NewMoney(resultAmount, to)
}
