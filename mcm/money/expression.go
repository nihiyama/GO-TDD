package money

// Expression is ...
type Expression interface {
	Plus(*Money) *Money
}
