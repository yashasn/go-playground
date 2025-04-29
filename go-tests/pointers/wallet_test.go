package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	//You don’t need to manually write &wallet.Deposit(10) — Go implicitly handles that
	wallet.Deposit(10)
	got := wallet.Balance()
	fmt.Printf("address of balance in test is %p \n", &wallet.balance)

	want := Bitcoin(0)
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
