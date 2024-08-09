package main

import (
	"fmt"
	"testing"
)

func TestPoint(t *testing.T) {

	wallet := Wallet{}
	wallet.Deposit(100)
	got := wallet.Balance()

	want := 100

	fmt.Printf("address of balance in test is %p \n", &wallet.balance)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
