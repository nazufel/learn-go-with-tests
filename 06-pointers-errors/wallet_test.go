package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

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

}
