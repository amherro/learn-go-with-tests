package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10

	fmt.Printf("The balance in the test is %p.\n", &wallet.balance)

	if got != want {
		t.Errorf("Got %d want %d", got, want)
	}
}
