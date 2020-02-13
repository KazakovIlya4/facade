package wallet

import (
	"fmt"
)

var (
	errInsufficientFunds = fmt.Errorf("insufficient funds")
)

// Wallet allows to check balance and withdraw money
type Wallet interface {
	Withdraw(amount uint32) (err error)
	Balance() (balance int)
}

type walletService struct {
	name    string
	balance int
}

// Withdraw decreases amount of money in walletService
func (w *walletService) Withdraw(amount uint32) (err error) {
	if w.balance < int(amount) {
		err = fmt.Errorf("user %s money deduction: %w", w.name, errInsufficientFunds)
		return
	}
	w.balance = w.balance - int(amount)
	return
}

func (w *walletService) Balance() (balance int) {
	balance = w.balance
	return
}

// NewWallet returns new instance of Wallet implementation
func NewWallet(name string, balance int) Wallet {
	return &walletService{
		name:    name,
		balance: balance,
	}
}
