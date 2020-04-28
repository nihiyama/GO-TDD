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

func TestMoneyStringfy(t *testing.T) {
	doller := NewDoller(5)
	actual := doller.ToString()
	expect := "[5, USD]"

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

func TestSumStringfy(t *testing.T) {
	doller := NewDoller(5)
	franc := NewFranc(5)
	sum := NewSum(*doller, *franc)
	actual := sum.ToString()
	expect := "[{5 USD}, {5 CHF}]"

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
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
	five := NewDoller(5)
	sum := five.plus(five)
	bank := NewBank()
	actual := *bank.reduce(sum, "USD")
	expect := *NewDoller(10)

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}

}

func TestPlusReturnSum(t *testing.T) {
	five := NewDoller(5)
	sum := five.plus(five)
	actual := sum.augend
	expect := *five

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}

	actual = sum.addend

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

func TestMoneyGetMoney(t *testing.T) {
	five := NewDoller(5)
	actualSlice := five.GetMoney()
	expectSlice := []Money{*NewDoller(5)}
	var actual, expect Money
	if len(expectSlice) == len(actualSlice) {
		for i := 0; i < len(actualSlice); i++ {
			actual = actualSlice[i]
			expect = expectSlice[i]
			if expect != actual {
				t.Errorf("%v != %v", expect, actual)
			}
		}
	} else {
		t.Errorf("Size of slice is not match(%d, %d)", len(expectSlice), len(actualSlice))
	}
}

func TestSumGetMoney(t *testing.T) {
	five := NewDoller(5)
	six := NewDoller(5)
	sum := NewSum(*five, *six)
	actualSlice := sum.GetMoney()
	expectSlice := []Money{*NewDoller(5), *NewDoller(5)}
	var actual, expect Money

	if len(expectSlice) == len(actualSlice) {
		for i := 0; i < len(actualSlice); i++ {
			actual = actualSlice[i]
			expect = expectSlice[i]
			if expect != actual {
				t.Errorf("%v != %v", expect, actual)
			}
		}
	} else {
		t.Errorf("Size of slice is not match(%d, %d)", len(expectSlice), len(actualSlice))
	}
}

func TestMoneyReduce(t *testing.T) {
	five := NewDoller(5)
	bank := NewBank()
	actual := *five.Reduce(bank, "USD")
	expect := *NewDoller(5)

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}
}

func TestSumReduce(t *testing.T) {
	sum := NewSum(*NewDoller(5), *NewDoller(5))
	bank := NewBank()
	actual := *sum.Reduce(bank, "USD")
	expect := *NewDoller(10)

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}
}

func TestBankReduce(t *testing.T) {
	sum := NewSum(*NewMoney(3), *NewMoney(4))
	bank := NewBank()
	actual := *bank.reduce(sum, "USD")
	expect := *NewDoller(7)

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}

	doller := NewDoller(7)
	actual = *bank.reduce(doller, "USD")

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}

}

func TestPairEquals(t *testing.T) {
	pair1 := NewPair("CHF", "USD")
	pair2 := NewPair("CHF", "USD")

	actual := pair1.equals(pair2)
	expect := true

	if expect != actual {
		t.Errorf("%t != %t", expect, actual)
	}

	pair3 := NewPair("YEN", "USD")
	actual = pair1.equals(pair3)
	expect = false

	if expect != actual {
		t.Errorf("%t != %t", expect, actual)
	}
}

func TestPairHashCode(t *testing.T) {
	pair1 := NewPair("CHF", "USD")
	actual := pair1.hachCode()
	expect := 0

	if expect != actual {
		t.Errorf("%d != %d", expect, actual)
	}
}

func TestReduceMoneyDifferentCurrency(t *testing.T) {
	bank := NewBank()
	bank.addRate("CHF", "USD", 2)
	actual := *bank.reduce(NewFranc(2), "USD")
	expect := *NewDoller(1)

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}
}
