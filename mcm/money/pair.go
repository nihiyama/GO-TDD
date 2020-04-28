package money

// Pair is ...
type Pair struct {
	from string
	to   string
}

// NewPair is ...
func NewPair(from string, to string) *Pair {
	return &Pair{from, to}
}

func (pair *Pair) equals(anotherPair *Pair) bool {
	return pair.from == anotherPair.from &&
		pair.to == anotherPair.to
}

func (pair *Pair) hachCode() int {
	return 0
}
