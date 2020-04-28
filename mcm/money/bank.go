package money

// Bank is ...
type Bank struct {
	rates map[Pair]int
}

// NewBank is ...
func NewBank() *Bank {
	rates := make(map[Pair]int, 20)
	return &Bank{rates}
}

func (bank *Bank) reduce(expression Expression, to string) *Money {
	return expression.Reduce(bank, to)
}

func (bank *Bank) addRate(from string, to string, rate int) {
	bank.rates[*NewPair(from, to)] = rate
}

func (bank *Bank) rate(from string, to string) int {
	if from == to {
		return 1
	}
	return bank.rates[*NewPair(from, to)]
}
