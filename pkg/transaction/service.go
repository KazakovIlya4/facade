package transaction

import (
	"facade/pkg/models"
	"sync"
)

// Transaction
type Transaction interface {
	Save(userID string, operation int, amount int, success bool) (id int)
	GetLast(userID string, limit int) (transactions []models.Record)
}

type transactionService struct {
	RWLock       *sync.RWMutex
	transactions []*models.Record
}

func (t *transactionService) Save(userID string, operation int, amount int, success bool) (id int) {
	t.RWLock.Lock()
	defer t.RWLock.Unlock()
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

func (t *transactionService) GetLast(userID string, limit int) (transactions []models.Record) {
	t.RWLock.RLock()
	defer t.RWLock.RUnlock()

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

// NewTransactionService returns new instance of transactionService implementation
func NewTransactionService() Transaction {
	return &transactionService{
		RWLock:       &sync.RWMutex{},
		transactions: make([]*models.Record, 0, 1024),
	}
}
