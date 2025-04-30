package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			//fatal stops ends the test and stops further assertions
			t.Fatal("wanted an error but didn't get one")
		}

		if got.Error() != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		//You don’t need to manually write &wallet.Deposit(10) — Go implicitly handles that
		wallet.Deposit(Bitcoin(10))
		fmt.Printf("address of balance in test is %p \n", &wallet.balance)

		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		//able to set balance because it is set as private within the same package
		wallet := Wallet{balance: Bitcoin(20)}
		//You don’t need to manually write &wallet.Deposit(10) — Go implicitly handles that
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		//able to set balance because it is set as private within the same package
		wallet := Wallet{balance: Bitcoin(20)}
		//You don’t need to manually write &wallet.Deposit(10) — Go implicitly handles that
		err := wallet.Withdraw(Bitcoin(30))

		assertError(t, err, ErrInsufficientFunds.Error())
		//assertBalance(t, wallet, Bitcoin(10))

	})

}
