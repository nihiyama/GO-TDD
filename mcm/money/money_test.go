package money

import (
	"testing"
)

func TestDollerMultiplication(t *testing.T) {
	doller := NewDoller(5)
	doller.times(2)
	actual := doller.Amount
	expect := 10

	if expect != actual {
		t.Errorf("%d != %d", expect, actual)
	}
}
