package wallet

import "github.com/stretchr/testify/mock"

type WalletMock struct {
	mock.Mock
}

func (w *WalletMock) Withdraw(amount uint32) (err error) {
	args := w.Called(amount)
	if a, ok := args.Get(0).(error); ok {
		err = a
	}
	return
}

func (w *WalletMock) Balance() (balance int) {
	args := w.Called()
	if a, ok := args.Get(0).(int); ok {
		balance = a
	}
	return
}
