package money

import (
	"reflect"
	"testing"
)

func TestMoneyGetCurrencyType(t *testing.T) {
	money := Money{5}
	actual := reflect.TypeOf(money.GetCurrency())
	expect := reflect.TypeOf(Money{5})

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}
}

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

func TestDollerGetCurrencyType(t *testing.T) {
	doller := NewDoller(5)
	actual := reflect.TypeOf(doller.GetCurrency())
	expect := reflect.TypeOf(*NewDoller(5))

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}
}

func TestDollerEquality(t *testing.T) {
	dollerFive1 := NewDoller(5)
	dollerFive2 := NewDoller(5)
	actual := Equals(dollerFive1, dollerFive2)
	expect := true

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}

	dollerFive2 = NewDoller(6)
	actual = Equals(dollerFive1, dollerFive2)
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

func TestFrancEquality(t *testing.T) {
	francFive1 := NewFranc(5)
	francFive2 := NewFranc(5)
	actual := Equals(francFive1, francFive2)
	expect := true

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}

	francFive2 = NewFranc(6)
	actual = Equals(francFive1, francFive2)
	expect = false

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}
}

func TestFrancGetCurrencyType(t *testing.T) {
	doller := NewFranc(5)
	actual := reflect.TypeOf(doller.GetCurrency())
	expect := reflect.TypeOf(*NewFranc(5))

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}
}
func TestFrancDollerEquality(t *testing.T) {
	francFive1 := NewFranc(5)
	dollerFive1 := NewDoller(5)
	actual := Equals(francFive1, dollerFive1)
	expect := false

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}
}
