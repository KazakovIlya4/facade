package transaction

import (
	"github.com/stretchr/testify/mock"

	"facade/pkg/models"
)

// Mock of transaction service
type Mock struct {
	mock.Mock
}

// Save ...
func (m *Mock) Save(userID string, operation int, amount int, success bool) (id int) {
	args := m.Called(userID, operation, amount, success)
	if a, ok := args.Get(0).(int); ok {
		id = a
	}
	return
}

// GetLast ...
func (m *Mock) GetLast(userID string, limit int) (transactions []models.Record) {
	args := m.Called(userID, limit)
	if a, ok := args.Get(0).([]models.Record); ok {
		transactions = a
	}
	return
}
