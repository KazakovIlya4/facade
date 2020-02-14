package facade

import (
	"fmt"
	"sync"

	"facade/pkg/models"
)

const (
	withdrawalCodeTest = iota
	balanceCodeTest
)

var (
	errInsufficientFunds = fmt.Errorf("insufficient funds")
	errAccountNotFound   = fmt.Errorf("account not found")
)

type wallet = interface {
	Withdraw(amount uint32) (err error)
	Balance() (balance int)
}

type transactionSystem interface {
	Save(userID string, operation int, amount int, success bool) (id int)
	GetLast(userID string, limit int) (transactions []models.Record)
}

// PaymentSystem represents money processor
// Withdraw withdraws `amount` funds from `id` account if `securityCode` corresponding to `transactionID` is correct
// Balance return balance  of `id` account if `securityCode` corresponding to `transactionID` is correct
type PaymentSystem interface {
	Withdraw(id string, amount uint32) (err error)
	Balance(id string) (balance int, err error)
}

type paymentSystemService struct {
	RWLock       *sync.RWMutex
	wallets      map[string]wallet
	transactions transactionSystem
}

func (p *paymentSystemService) Withdraw(userID string, amount uint32) (err error) {
	p.RWLock.RLock()
	defer p.RWLock.RUnlock()
	w, ok := p.wallets[userID]
	if ok != true {
		err = fmt.Errorf("get money %s: %w", userID, errAccountNotFound)
		return
	}

	success := true
	err = w.Withdraw(amount)
	if err != nil {
		err = fmt.Errorf("get money: %w", err)
		success = false
	}

	operationID := p.transactions.Save(userID, withdrawalCodeTest, int(amount), success)
	if !success {
		err = fmt.Errorf("operation id %d: %w", operationID, err)
		return
	}

	return
}

func (p *paymentSystemService) Balance(userID string) (balance int, err error) {
	p.RWLock.RLock()
	defer p.RWLock.RUnlock()
	w, ok := p.wallets[userID]
	if !ok {
		err = fmt.Errorf("get balance of %s: %w", userID, errAccountNotFound)
		return
	}

	balance = w.Balance()
	p.transactions.Save(userID, balanceCodeTest, balance, true)
	return
}

// NewPaymentSystem creates PaymentSystem implementation with wallets provided by `accounts`
func NewPaymentSystem(wallets map[string]wallet, transactions transactionSystem) PaymentSystem {
	return &paymentSystemService{
		RWLock:       &sync.RWMutex{},
		wallets:      wallets,
		transactions: transactions,
	}
}
