package money

// Doller ...
type Doller struct {
	Money
}

// NewDoller ...
func NewDoller(amount int) *Doller {
	return &Doller{Money{amount}}
}

func (doller *Doller) times(multiple int) *Doller {
	return &Doller{Money{doller.amount * multiple}}
}

func (doller *Doller) equals(anotherDoller *Doller) bool {
	return doller.amount == anotherDoller.amount
}
