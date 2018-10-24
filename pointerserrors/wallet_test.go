package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		initialBalance := Bitcoin(20)
		wallet := Wallet{balance: initialBalance}
		err := wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("withdraw with insufficient funds", func(t *testing.T) {
		initialBalance := Bitcoin(20)
		wallet := Wallet{balance: initialBalance}
		err := wallet.Withdraw(30)

		assertBalance(t, wallet, initialBalance)
		assertError(t, err, ErrInsufficientBalance)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("didn't get an error but wanted one")
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Errorf("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
