package facade

import (
	"fmt"
	"sync"

	"facade/pkg/models"
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
	sync.RWMutex
	wallets        map[string]wallet
	transactions   transactionSystem
	operationCodes map[string]int
}

func (p *paymentSystemService) Withdraw(userID string, amount uint32) (err error) {
	p.RLock()
	defer p.RUnlock()
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

	operationID := p.transactions.Save(userID, p.operationCodes["withdraw"], int(amount), success)
	if !success {
		err = fmt.Errorf("operation id %d: %w", operationID, err)
		return
	}

	return
}

func (p *paymentSystemService) Balance(userID string) (balance int, err error) {
	p.RLock()
	defer p.RUnlock()
	w, ok := p.wallets[userID]
	if !ok {
		err = fmt.Errorf("get balance of %s: %w", userID, errAccountNotFound)
		return
	}

	balance = w.Balance()
	p.transactions.Save(userID, p.operationCodes["balance"], balance, true)
	return
}

// NewPaymentSystem creates PaymentSystem implementation with wallets provided by `accounts`
func NewPaymentSystem(wallets map[string]wallet, transactions transactionSystem, codes map[string]int) PaymentSystem {
	return &paymentSystemService{
		wallets:        wallets,
		transactions:   transactions,
		operationCodes: codes,
	}
}
