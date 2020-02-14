package wallet

import "github.com/stretchr/testify/mock"

// Mock of wallet service
type Mock struct {
	mock.Mock
}

// Withdraw ...
func (m *Mock) Withdraw(amount uint32) (err error) {
	args := m.Called(amount)
	if a, ok := args.Get(0).(error); ok {
		err = a
	}
	return
}

// Balance ...
func (m *Mock) Balance() (balance int) {
	args := m.Called()
	if a, ok := args.Get(0).(int); ok {
		balance = a
	}
	return
}
