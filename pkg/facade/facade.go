package facade

import (
	"fmt"

	"facade/pkg/security"
	"facade/pkg/wallet"
)

var (
	insufficientFunds   = fmt.Errorf("insufficient funds")
	accountNotFound     = fmt.Errorf("account not found")
	invalidSecurityCode = fmt.Errorf("invalid securityChecker code")
)

// PaymentSystem represents money processor
// GetMoney withdraws `amount` funds from `id` account if `securityCode` corresponding to `transactionID` is correct
// GetBalance return balance  of `id` account if `securityCode` corresponding to `transactionID` is correct
type PaymentSystem interface {
	GetMoney(id string, amount uint32, securityCode int, transactionID int) error
	GetBalance(id string, sercurityCode int, transactionID int) (int, error)
}

type AccountInfo struct {
	ID     string
	wallet wallet.Wallet
}

type paymentSystemImplementation struct {
	wallets         map[string]wallet.Wallet
	securityChecker security.Checker
}

func (p *paymentSystemImplementation) GetMoney(userID string, amount uint32, code int, transactionID int) error {
	w, ok := p.wallets[userID]
	if ok != true {
		return fmt.Errorf("get money %s: %w", userID, accountNotFound)
	}
	success := p.securityChecker.Check(transactionID, code)
	if !success {
		return fmt.Errorf("get money %d: %w", code, invalidSecurityCode)
	}
	err := w.DeductMoney(amount)
	if err != nil {
		return fmt.Errorf("get money: %w", err)
	}
	return nil
}

func (p *paymentSystemImplementation) GetBalance(id string, code int, transactionID int) (int, error) {
	w, ok := p.wallets[id]
	if !ok {
		return 0, fmt.Errorf("get balance of %s: %w", id, accountNotFound)
	}
	success := p.securityChecker.Check(transactionID, code)
	if !success {
		return 0, fmt.Errorf("get balance of %s: %w", id, invalidSecurityCode)
	}
	return w.GetBalance(), nil
}

// NewPaymentSystem creates PaymentSystem implementation with wallets provided by `accounts`
func NewPaymentSystem(accounts []AccountInfo, checker security.Checker) PaymentSystem {
	wallets := make(map[string]wallet.Wallet, len(accounts))
	for _, v := range accounts {
		wallets[v.ID] = v.wallet
	}
	return &paymentSystemImplementation{
		wallets:         wallets,
		securityChecker: checker,
	}
}

func Account(id string, newWallet wallet.Wallet) (account AccountInfo) {
	account.ID = id
	account.wallet = newWallet
	return
}
