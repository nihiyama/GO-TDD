package money

// Expression is ...
type Expression interface {
	ToString() string
	GetMoney() []Money
	Reduce(*Bank, string) *Money
}
