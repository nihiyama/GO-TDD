package money

// Franc ...
type Franc struct {
	amount int
}

// NewFranc ...
func NewFranc(amount int) *Franc {
	return &Franc{amount}
}

func (franc *Franc) times(multiple int) *Franc {
	return &Franc{franc.amount * multiple}
}

func (franc *Franc) equals(anotherFranc *Franc) bool {
	return franc.amount == anotherFranc.amount
}
