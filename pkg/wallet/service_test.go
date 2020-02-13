package wallet

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWallet_WithdrawSuccess(t *testing.T) {
	wallet := NewWallet("Alice", 500)
	err := wallet.Withdraw(500)
	assert.Equal(t, nil, err)
	balance := wallet.Balance()
	assert.Equal(t, 0, balance)
}

func TestWallet_WithdrawInsufficientFundsFail(t *testing.T) {
	wallet := NewWallet("Alice", 500)
	err := wallet.Withdraw(501)
	assert.Equal(t, errors.As(err, &errInsufficientFunds), true)
	balance := wallet.Balance()
	assert.Equal(t, 500, balance)
}
