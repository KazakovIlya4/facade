package operation

import (
	"sync"

	"facade/pkg/models"
)

// Transaction saves operations info and retrieves them
type Transaction interface {
	Save(userID string, operation int, amount int, success bool) (id int)
	GetLast(userID string, limit int) (transactions []models.Record)
}

type service struct {
	sync.RWMutex
	transactions []*models.Record
}

func (t *service) Save(userID string, operation int, amount int, success bool) (id int) {
	t.Lock()
	defer t.Unlock()
	id = len(t.transactions)
	t.transactions = append(t.transactions, &models.Record{
		ID:            id,
		OperationCode: operation,
		UserID:        userID,
		Amount:        amount,
		Success:       success,
	})
	return
}

func (t *service) GetLast(userID string, limit int) (transactions []models.Record) {
	t.RLock()
	defer t.RUnlock()

	if limit < 0 {
		return
	}
	if limit == 0 {
		limit = len(t.transactions)
	}

	for i := len(t.transactions) - 1; i >= 0 && limit > 0; i-- {
		if t.transactions[i].UserID == userID {
			transactions = append(transactions, *t.transactions[i])
			limit--
		}
	}
	return
}

// NewService returns new instance of service implementation
func NewService() Transaction {
	return &service{
		transactions: make([]*models.Record, 0, 1024),
	}
}
