package facade

import (
	"fmt"
)

var (
	InsufficientFunds   = fmt.Errorf("insufficient funds")
	AccountNotFound     = fmt.Errorf("account not found")
	InvalidSecurityCode = fmt.Errorf("invalid security code")
)

const (
	securityCodeExampleValid int = 123
	transactionIDExample     int = 1234
)

type FacadePaymentSystem interface {
	GetMoney(id string, amount uint32, securityCode int, transactionID int) error
	GetBalance(id string, sercurityCode int, transactionID int) (int, error)
}

type accountInfo struct {
	name    string
	balance int
}

func CreatePaymentSystem(accounts []*accountInfo) FacadePaymentSystem {
	wallets := make(map[string]Wallet, len(accounts))
	for _, v := range accounts {
		wallets[v.name] = &wallet{
			name:    v.name,
			balance: v.balance,
		}
	}
	return &paymentSystem{
		wallets:  wallets,
		security: &securityCheck{},
	}
}

type paymentSystem struct {
	wallets  map[string]Wallet
	security SecurityCheck
}

func (s *paymentSystem) GetMoney(userID string, amount uint32, code int, transactionID int) error {
	w, ok := s.wallets[userID]
	if ok != true {
		return fmt.Errorf("get money %s: %w", userID, AccountNotFound)
	}
	success := s.security.check(transactionID, code)
	if !success {
		return fmt.Errorf("get money %d: %w", code, InvalidSecurityCode)
	}
	err := w.deductMoney(amount)
	if err != nil {
		return fmt.Errorf("get money: %w", err)
	}
	return nil
}
func (s *paymentSystem) GetBalance(id string, code int, transactionID int) (int, error) {
	w, ok := s.wallets[id]
	if !ok {
		return 0, fmt.Errorf("get balance of %s: %w", id, AccountNotFound)
	}
	success := s.security.check(transactionID, code)
	if !success {
		return 0, fmt.Errorf("get balance of %s: %w", id, InvalidSecurityCode)
	}
	return w.getBalance(), nil
}

type Wallet interface {
	deductMoney(amount uint32) error
	getBalance() int
}

type wallet struct {
	name    string
	balance int
}

func (s *wallet) deductMoney(amount uint32) error {
	fmt.Printf("Deducting %d from %s\n", amount, s.name)
	if s.balance < int(amount) {
		return fmt.Errorf("user %s money deduction: %w", s.name, InsufficientFunds)
	}
	s.balance = s.balance - int(amount)
	return nil
}

func (s wallet) getBalance() int {
	fmt.Printf("Getting balance of %s\n", s.name)
	return s.balance
}

type SecurityCheck interface {
	check(transactionID int, code int) bool
}

type securityCheck struct{}

func (s securityCheck) check(_ int, code int) bool {
	fmt.Printf("Checking security code %d\n", code)
	//placeholder
	if code == securityCodeExampleValid {
		return true
	}
	return false
}
