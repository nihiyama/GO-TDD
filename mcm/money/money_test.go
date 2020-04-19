package money

import (
	"testing"
)

func TestDollerMultiplication(t *testing.T) {
	doller := NewDoller(5)
	actual := *(doller.times(2))
	expect := *(NewDoller(10))

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}

	actual = *(doller.times(3))
	expect = *(NewDoller(15))

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
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

func TestFrancMultiplication(t *testing.T) {
	franc := NewFranc(5)
	actual := *(franc.times(2))
	expect := *(NewFranc(10))

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}

	actual = *(franc.times(3))
	expect = *(NewFranc(15))

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}
}
