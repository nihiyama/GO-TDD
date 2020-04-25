package money

import (
	"testing"
)

// func TestMoneyGetCurrencyType(t *testing.T) {
// 	money := NewMoney(5)
// 	actual := reflect.TypeOf(money.GetCurrency())
// 	expect := reflect.TypeOf(*NewMoney(5))

// 	if expect != actual {
// 		t.Errorf("%v != %v", expect, actual)
// 	}
// }

func TestMoneyMultiplication(t *testing.T) {
	money := NewMoney(5)
	actual := *money.times(2)
	expect := *NewMoney(10)

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}
}

func TestMoneyEquality(t *testing.T) {
	moneyFive1 := NewMoney(5)
	moneyFive2 := NewMoney(5)
	actual := moneyFive1.equals(moneyFive2)
	expect := true

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}

	moneyFive2 = NewMoney(6)
	actual = moneyFive1.equals(moneyFive2)
	expect = false

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}
}

func TestSimpleAddition(t *testing.T) {
	money := NewMoney(5)
	actual := *money.plus(NewMoney(5))
	expect := *NewMoney(10)

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

	dollerFive2 = NewDoller(6)
	actual = dollerFive1.equals(dollerFive2)
	expect = false

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}

}

func TestFrancEquality(t *testing.T) {
	francFive1 := NewFranc(5)
	francFive2 := NewFranc(5)
	actual := francFive1.equals(francFive2)
	expect := true

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}

	francFive2 = NewFranc(6)
	actual = francFive1.equals(francFive2)
	expect = false

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}
}

func TestFrancDollerEquality(t *testing.T) {
	francFive1 := NewFranc(5)
	dollerFive1 := NewDoller(5)
	actual := francFive1.equals(dollerFive1)
	expect := false

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}

	moneyFive1 := NewMoney(5, "CHF")
	actual = francFive1.equals(moneyFive1)
	expect = true

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}

	actual = dollerFive1.equals(moneyFive1)
	expect = false

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}

}
