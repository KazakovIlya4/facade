package transaction

import (
	"github.com/stretchr/testify/mock"

	"facade/pkg/models"
)

type Mock struct {
	mock.Mock
}

func (w *Mock) Save(userID string, operation int, amount int, success bool) (id int) {
	args := w.Called(userID, operation, amount, success)
	if a, ok := args.Get(0).(int); ok {
		id = a
	}
	return
}

func (w *Mock) GetLast(userID string, limit int) (transactions []models.Record) {
	args := w.Called(userID, limit)
	if a, ok := args.Get(0).([]models.Record); ok {
		transactions = a
	}
	return
}
