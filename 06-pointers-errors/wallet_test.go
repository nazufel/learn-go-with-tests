package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("depost", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Depost(Bitcoin(10))

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})

}
