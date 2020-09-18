package change

import (
	"testing"
)

func TestChange1000MustHave1BankNoteOf1000(t *testing.T) {
	bcArray := Change(1000.00)
	if bcArray[0].Amount != 1 {
		t.Error("Must 1 Have Bank 1000 but have", bcArray[0].Amount)
	}
}
func TestChange1688point75MustHave1OfEachBankOrCoin(t *testing.T) {
	bcArray := Change(1688.75)
	if bcArray[0].Amount != 1 {
		t.Error("Must Have 1 Bank 1000 but have", bcArray[0].Amount)
	}
	if bcArray[1].Amount != 1 {
		t.Error("Must 1 Have Bank 500 but have", bcArray[1].Amount)
	}
	if bcArray[2].Amount != 1 {
		t.Error("Must 1 Have Bank 100 but have", bcArray[2].Amount)
	}
	if bcArray[3].Amount != 1 {
		t.Error("Must 1 Have Bank 50 but have", bcArray[3].Amount)
	}
	if bcArray[4].Amount != 1 {
		t.Error("Must 1 Have Bank 20 but have", bcArray[4].Amount)
	}
	if bcArray[5].Amount != 1 {
		t.Error("Must 1 Have Coin 10 but have", bcArray[5].Amount)
	}
	if bcArray[6].Amount != 1 {
		t.Error("Must 1 Have Coin 5 but have", bcArray[6].Amount)
	}
	if bcArray[7].Amount != 1 {
		t.Error("Must 1 Have Coin 2 but have", bcArray[7].Amount)
	}
	if bcArray[8].Amount != 1 {
		t.Error("Must 1 Have Coin 1 but have", bcArray[8].Amount)
	}
	if bcArray[9].Amount != 1 {
		t.Error("Must 1 Have Coin 0.5 but have", bcArray[9].Amount)
	}
	if bcArray[10].Amount != 1 {
		t.Error("Must 1 Have Coin 0.25 but have", bcArray[10].Amount)
	}
}
