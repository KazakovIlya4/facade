package facade

import (
	"fmt"
	"sync"
)

var (
	errInsufficientFunds   = fmt.Errorf("insufficient funds")
	errAccountNotFound     = fmt.Errorf("account not found")
	errInvalidSecurityCode = fmt.Errorf("invalid securityChecker code")
)

type wallet interface {
	Withdraw(amount uint32) (err error)
	Balance() (balance int)
}

type checker interface {
	Check(securityCode, transactionID int) bool
}

// PaymentSystem represents money processor
// Withdraw withdraws `amount` funds from `id` account if `securityCode` corresponding to `transactionID` is correct
// Balance return balance  of `id` account if `securityCode` corresponding to `transactionID` is correct
type PaymentSystem interface {
	Withdraw(id string, amount uint32, securityCode int, transactionID int) (err error)
	Balance(id string, securityCode int, transactionID int) (balance int, err error)
}

type paymentSystemService struct {
	RWLock          *sync.RWMutex
	wallets         map[string]wallet
	securityChecker checker
}

func (p *paymentSystemService) Withdraw(userID string, amount uint32, securityCode int, transactionID int) (err error) {
	p.RWLock.RLock()
	defer p.RWLock.RUnlock()
	w, ok := p.wallets[userID]
	if ok != true {
		err = fmt.Errorf("get money %s: %w", userID, errAccountNotFound)
		return
	}
	success := p.securityChecker.Check(securityCode, transactionID)
	if !success {
		err = fmt.Errorf("get money %d: %w", securityCode, errInvalidSecurityCode)
		return
	}
	err = w.Withdraw(amount)
	if err != nil {
		err = fmt.Errorf("get money: %w", err)
	}
	return
}

func (p *paymentSystemService) Balance(id string, code int, transactionID int) (balance int, err error) {
	p.RWLock.RLock()
	defer p.RWLock.RUnlock()
	w, ok := p.wallets[id]
	if !ok {
		err = fmt.Errorf("get balance of %s: %w", id, errAccountNotFound)
		return
	}
	success := p.securityChecker.Check(transactionID, code)
	if !success {
		err = fmt.Errorf("get balance of %s: %w", id, errInvalidSecurityCode)
		return
	}

	balance = w.Balance()
	return
}

// NewPaymentSystem creates PaymentSystem implementation with wallets provided by `accounts`
func NewPaymentSystem(wallets map[string]wallet, checker checker) PaymentSystem {
	return &paymentSystemService{
		RWLock:          &sync.RWMutex{},
		wallets:         wallets,
		securityChecker: checker,
	}
}
