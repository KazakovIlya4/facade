package facade

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	wallet_mock "facade/pkg/account"
	"facade/pkg/operation"
)

const (
	withdrawalCode = iota
	balanceCode
	transactionServiceMethodName           = "Save"
	walletServiceMethodWithdrawName        = "Withdraw"
	walletServiceMethodBalanceName         = "Balance"
	withdrawAmount500Test           uint32 = 500
	withdrawAmountTooMuchTest       uint32 = 2 * withdrawAmount500Test
)

func setupTests() (paymentSystem PaymentSystem) {
	transactionServiceMock := &operation.Mock{
		Mock: mock.Mock{},
	}
	transactionServiceMock.On(transactionServiceMethodName, mock.AnythingOfType("string"),
		mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("bool")).
		Return(0)

	transactionServiceMock.On(walletServiceMethodBalanceName, mock.AnythingOfType("string"),
		mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("bool")).
		Return(0)

	serviceWalletMock := &wallet_mock.Mock{
		Mock: mock.Mock{},
	}
	serviceWalletMock.On(walletServiceMethodWithdrawName, withdrawAmount500Test).
		Return(nil)
	serviceWalletMock.On(walletServiceMethodWithdrawName, withdrawAmountTooMuchTest).
		Return(errInsufficientFunds)

	paymentSystem = NewPaymentSystem(map[string]wallet{"Alice": serviceWalletMock},
		transactionServiceMock, map[string]int{"withdraw": withdrawalCode, "balance": balanceCode})

	return
}

func Test_WithdrawSuccess(t *testing.T) {
	paymentSystem := setupTests()
	err := paymentSystem.Withdraw("Alice", withdrawAmount500Test)
	assert.Equal(t, nil, err)
}

func Test_WithdrawNotFoundFail(t *testing.T) {
	paymentSystem := setupTests()
	err := paymentSystem.Withdraw("Bob", withdrawAmount500Test)
	assert.Equal(t, true, errors.As(err, &errAccountNotFound))
}

func Test_WithdrawInsufficientFundsFail(t *testing.T) {
	paymentSystem := setupTests()
	err := paymentSystem.Withdraw("Alice", withdrawAmountTooMuchTest)
	assert.Equal(t, true, errors.As(err, &errInsufficientFunds))
}
