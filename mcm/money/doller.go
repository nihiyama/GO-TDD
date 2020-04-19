package money

// Doller ...
type Doller struct {
	Amount int
}

// NewDoller ...
func NewDoller(amount int) *Doller {
	return &Doller{amount}
}

func (doller *Doller) times(multiple int) *Doller {
	doller.Amount *= 2
	return doller
}
