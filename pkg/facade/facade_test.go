package facade

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"facade/pkg/transaction"
	wallet_mock "facade/pkg/wallet"
)

const (
	transactionServiceMethodName           = "Save"
	walletServiceMethodWithdrawName        = "Withdraw"
	walletServiceMethodBalanceName         = "Balance"
	withdrawAmount500Test           uint32 = 500
	withdrawAmountTooMuchTest       uint32 = 2 * withdrawAmount500Test
)

func setupTests() (paymentSystem PaymentSystem) {
	transactionServiceMock := &transaction.Mock{
		Mock: mock.Mock{},
	}
	transactionServiceMock.On(transactionServiceMethodName, mock.AnythingOfType("string"),
		mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("bool")).
		Return(0)

	transactionServiceMock.On(walletServiceMethodBalanceName, mock.AnythingOfType("string"),
		mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("bool")).
		Return(0)

	serviceWalletMock := &wallet_mock.WalletMock{
		Mock: mock.Mock{},
	}
	serviceWalletMock.On(walletServiceMethodWithdrawName, withdrawAmount500Test).
		Return(nil)
	serviceWalletMock.On(walletServiceMethodWithdrawName, withdrawAmountTooMuchTest).
		Return(errInsufficientFunds)

	paymentSystem = NewPaymentSystem(map[string]wallet{"Alice": serviceWalletMock},
		transactionServiceMock)

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
