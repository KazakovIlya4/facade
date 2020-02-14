package facade

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

const (
	checkerServiceMethodName               = "Check"
	walletServiceMethodWithdrawName        = "Withdraw"
	walletServiceMethodBalanceName         = "Balance"
	withdrawAmount500Test           uint32 = 500
	withdrawAmountTooMuchTest       uint32 = 2 * withdrawAmount500Test
	testSecurityCodeValid           int    = 123
	testSecurityCodeInvalid         int    = 321
	testTransactionID               int    = 1234
)

func setupTests() (paymentSystem PaymentSystem) {
	securityCheckerMock := &CheckerMock{
		Mock: mock.Mock{},
	}
	securityCheckerMock.On(checkerServiceMethodName, testSecurityCodeValid, testTransactionID).
		Return(true)
	securityCheckerMock.On(checkerServiceMethodName, testSecurityCodeInvalid, testTransactionID).
		Return(false)

	serviceWalletMock := &WalletMock{
		name:    "Alice",
		balance: int(withdrawAmount500Test) + 100,
		Mock:    mock.Mock{},
	}
	serviceWalletMock.On(walletServiceMethodWithdrawName, withdrawAmount500Test).
		Return(nil)
	serviceWalletMock.On(walletServiceMethodWithdrawName, withdrawAmountTooMuchTest).
		Return(errInsufficientFunds)

	paymentSystem = NewPaymentSystem(map[string]wallet{"Alice": serviceWalletMock},
		securityCheckerMock)

	return
}

type CheckerMock struct {
	mock.Mock
}

func (c *CheckerMock) Check(securityCode, transactionID int) (valid bool) {
	args := c.Called(securityCode, transactionID)
	if a, ok := args.Get(0).(bool); ok {
		valid = a
	}
	return
}

func Test_WithdrawSuccess(t *testing.T) {
	paymentSystem := setupTests()
	err := paymentSystem.Withdraw("Alice", withdrawAmount500Test, testSecurityCodeValid, testTransactionID)
	assert.Equal(t, nil, err)
}

func Test_WithdrawNotFoundFail(t *testing.T) {
	paymentSystem := setupTests()
	err := paymentSystem.Withdraw("Bob", withdrawAmount500Test, testSecurityCodeValid, testTransactionID)
	assert.Equal(t, true, errors.As(err, &errAccountNotFound))
}

func Test_WithdrawCheckFail(t *testing.T) {
	paymentSystem := setupTests()
	err := paymentSystem.Withdraw("Alice", withdrawAmount500Test, testSecurityCodeInvalid, testTransactionID)
	assert.Equal(t, true, errors.As(err, &errInvalidSecurityCode))
}

func Test_WithdrawInsufficientFundsFail(t *testing.T) {
	paymentSystem := setupTests()
	err := paymentSystem.Withdraw("Alice", withdrawAmountTooMuchTest, testSecurityCodeValid, testTransactionID)
	assert.Equal(t, true, errors.As(err, &errInsufficientFunds))
}
