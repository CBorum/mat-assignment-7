package main

import (
	"fmt"

	contract "github.com/cborum/design-by-contract"
)

// Account ...
type Account struct {
	balance float64
}

func main() {
	account := &Account{
		balance: 100,
	}
	account.Deposit(100)
	fmt.Println("balance", account.balance)
	// account.Deposit(-100)
	// fmt.Println("balance", account.balance)
	account.Withdraw(100)
	fmt.Println("balane", account.balance)
	account.Withdraw(100)
	fmt.Println("balance", account.balance)
}

// Deposit ...
func (a *Account) Deposit(amount float64) {
	contract.IsPositive(amount).Check()
	a.balance += amount
}

// Withdraw ...
func (a *Account) Withdraw(amount float64) {
	contract.Wrap(
		contract.IsPositive(amount),
		contract.IsZeroOrPositive(a.balance-amount),
	)
	a.balance -= amount
}
