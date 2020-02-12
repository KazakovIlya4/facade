package facade

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"facade/pkg/security"
	"facade/pkg/wallet"
)

const (
	securityCodeExampleValid int = 123
	transactionIDExample     int = 1234
)

func TestGetMoney(t *testing.T) {
	paymentSystem := NewPaymentSystem([]AccountInfo{
		Account("1", wallet.NewWallet("1", 500)),
		Account("2", wallet.NewWallet("2", 50)),
	}, security.NewChecker())

	err := paymentSystem.GetMoney("1", 450, securityCodeExampleValid, transactionIDExample)
	assert.Equal(t, nil, err)
	balance, err := paymentSystem.GetBalance("1", securityCodeExampleValid, transactionIDExample)
	assert.Equal(t, nil, err)
	assert.Equal(t, 50, balance)

	err = paymentSystem.GetMoney("2", 450, 12, transactionIDExample)
	assert.Equal(t, true, errors.As(err, &invalidSecurityCode))

	err = paymentSystem.GetMoney("2", 450, securityCodeExampleValid, transactionIDExample)
	assert.Equal(t, true, errors.As(err, &insufficientFunds))

	balance, err = paymentSystem.GetBalance("2", securityCodeExampleValid, transactionIDExample)
	assert.Equal(t, nil, err)
	assert.Equal(t, 50, balance)
}
