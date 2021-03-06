package main

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()
	want := Bitcoin(22)

	if got != want {
		t.Errorf("got %s and want %s", got, want)
	}

}
