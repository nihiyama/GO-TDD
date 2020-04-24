package money

import (
	"reflect"
	"testing"
)

func TestMoneyGetCurrencyType(t *testing.T) {
	money := NewMoney(5)
	actual := reflect.TypeOf(money.GetCurrency())
	expect := reflect.TypeOf(*NewMoney(5))

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}
}

func TestMoneyMultiplication(t *testing.T) {
	money := NewMoney(5)
	actual := *money.Times(2)
	expect := *NewMoney(10)

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}
}

func TestMoneyCurrencyName(t *testing.T) {
	money := NewMoney(5)
	actual := money.GetCurrencyName()
	expect := "None"

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}

	doller := NewDoller(5)
	actual = doller.GetCurrencyName()
	expect = "USD"

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}

	franc := NewFranc(5)
	actual = franc.GetCurrencyName()
	expect = "CHF"

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}

}

func TestMoneyEquality(t *testing.T) {
	moneyFive1 := NewMoney(5)
	moneyFive2 := NewMoney(5)
	actual := Equals(moneyFive1, moneyFive2)
	expect := true

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}

	moneyFive2 = NewMoney(6)
	actual = Equals(moneyFive1, moneyFive2)
	expect = false

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

func TestFrancDollerEquality(t *testing.T) {
	francFive1 := NewFranc(5)
	dollerFive1 := NewDoller(5)
	actual := Equals(francFive1, dollerFive1)
	expect := false

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}

	moneyFive1 := NewMoney(5, "CHF")
	actual = Equals(francFive1, moneyFive1)
	expect = true

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}

	actual = Equals(dollerFive1, moneyFive1)
	expect = false

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}

}
