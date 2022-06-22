package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertError := func(t testing.TB, err error, want string) {
		t.Helper()
		if err == nil {
			t.Fatal("wanted an error but didn't get one")
		}

		if err.Error() != want {
			t.Errorf("got: %q, want: %q", err, want)
		}
	}
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {

		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	}

	t.Run("depost", func(t *testing.T) {

		wallet := Wallet{}
		wallet.Depost(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("withdraw", func(t *testing.T) {

		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, "cannot withdraw, insufficient funds")
		assertBalance(t, wallet, startingBalance)
	})

}
