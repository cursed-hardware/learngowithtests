package pointers

import (
	"errors"
	"testing"
)

func TestBitcoinStringer(t *testing.T) {
	btc := Bitcoin(10)
	got := btc.String()
	want := "10 BTC"
	if got != want {
		t.Errorf("got %s but want %s", got, want)
	}
}

func TestWallet(t *testing.T) {
	assertMsg := func(t testing.TB, got, want Bitcoin) {
		t.Helper()
		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	}

	assertErr := func(t testing.TB, got, want error) {
		t.Helper()

		if got == nil {
			t.Error("expected error but didn't get one")
		}

		if !errors.Is(got, want) {
			t.Errorf("got %q but want %q", got.Error(), want.Error())
		}
	}

	assertNoErr := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Errorf("expected no error but got %q", got.Error())
		}
	}

	t.Run("deposit correct amount of btc", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertMsg(t, wallet.Balance(), Bitcoin(10))
	})

	t.Run("withdraw correct amount of btc", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertMsg(t, wallet.Balance(), Bitcoin(10))
		assertNoErr(t, err)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(100))
		assertMsg(t, wallet.Balance(), Bitcoin(10))
		assertErr(t, err, ErrInsufficientBalance)
	})
}
