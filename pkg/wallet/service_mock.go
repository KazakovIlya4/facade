package wallet

import "github.com/stretchr/testify/mock"

type Mock struct {
	name    string
	balance int
	mock.Mock
}

func (w *Mock) Withdraw(amount uint32) (err error) {
	args := w.Called(amount)
	if a, ok := args.Get(0).(error); ok {
		err = a
	}
	return
}

func (w *Mock) Balance() (balance int) {
	args := w.Called()
	if a, ok := args.Get(0).(int); ok {
		balance = a
	}
	return
}
