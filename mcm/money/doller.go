package money

// Doller ...
type Doller struct {
	amount int
}

// NewDoller ...
func NewDoller(amount int) *Doller {
	return &Doller{amount}
}

func (doller *Doller) times(multiple int) *Doller {
	return &Doller{doller.amount * multiple}
}

func (doller *Doller) equals(anotherDoller *Doller) bool {
	return doller.amount == anotherDoller.amount
}
