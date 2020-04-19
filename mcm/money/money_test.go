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

func TestDollerEquality(t *testing.T) {
	dollerFive1 := NewDoller(5)
	dollerFive2 := NewDoller(5)
	actual := dollerFive1.equals(dollerFive2)
	expect := true

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}

	dollerFive3 := NewDoller(6)
	actual = dollerFive1.equals(dollerFive3)
	expect = false

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}

}
