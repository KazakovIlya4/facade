package facade

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMoney(t *testing.T) {
	paymentSystem := CreatePaymentSystem([]*accountInfo{
		{name: "1", balance: 500},
		{name: "2", balance: 50},
	})

	err := paymentSystem.GetMoney("1", 450, securityCodeExampleValid, transactionIDExample)
	assert.Equal(t, nil, err)пше
	balance, err := paymentSystem.GetBalance("1", securityCodeExampleValid, transactionIDExample)
	assert.Equal(t, nil, err)
	assert.Equal(t, 50, balance)

	err = paymentSystem.GetMoney("2", 450, 12, transactionIDExample)
	assert.Equal(t, true, errors.As(err, &InvalidSecurityCode))

	err = paymentSystem.GetMoney("2", 450, securityCodeExampleValid, transactionIDExample)
	assert.Equal(t, true, errors.As(err, &InsufficientFunds))

	balance, err = paymentSystem.GetBalance("2", securityCodeExampleValid, transactionIDExample)
	assert.Equal(t, nil, err)
	assert.Equal(t, 50, balance)
}
