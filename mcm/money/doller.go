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
	return &Doller{doller.Amount * multiple}
}

func (doller *Doller) equals(anotherDoller *Doller) bool {
	return doller.Amount == anotherDoller.Amount
}
