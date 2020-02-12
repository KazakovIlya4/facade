package wallet

import "fmt"

var (
	insufficientFunds = fmt.Errorf("insufficient funds")
)

type wallet interface {
	DeductMoney(amount uint32) error
	GetBalance() int
}

type Wallet interface {
	DeductMoney(amount uint32) error
	GetBalance() int
}

type walletImplementation struct {
	name    string
	balance int
}

// DeductMoney decreases amount of money in walletImplementation
func (w *walletImplementation) DeductMoney(amount uint32) error {
	fmt.Printf("Deducting %d from %s\n", amount, w.name)
	if w.balance < int(amount) {
		return fmt.Errorf("user %s money deduction: %w", w.name, insufficientFunds)
	}
	w.balance = w.balance - int(amount)
	return nil
}

func (w *walletImplementation) GetBalance() int {
	fmt.Printf("Getting balance of %s\n", w.name)
	return w.balance
}

// NewWallet returns new instance of Wallet implementation
func NewWallet(name string, balance int) Wallet {
	return &walletImplementation{
		name:    name,
		balance: balance,
	}
}
