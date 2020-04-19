package money

import (
	"testing"
)

func TestDollerMultiplication(t *testing.T) {
	doller := NewDoller(5)
	product := doller.times(2)
	actual := product.Amount
	expect := 10

	if expect != actual {
		t.Errorf("%d != %d", expect, actual)
	}

	product = doller.times(3)
	actual = product.Amount
	expect = 15

	if expect != actual {
		t.Errorf("%d != %d", expect, actual)
	}

}
