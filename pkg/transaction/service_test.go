package transaction

import (
	"facade/pkg/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	withdrawalCodeTest = iota
	balanceCodeTest
)

func TestTransaction_Get(t *testing.T) {

	transactions := NewTransactionService()
	transactions.Save("Alice", withdrawalCodeTest, 400, true)
	rec := transactions.GetLast("Bob", 0)
	assert.Equal(t, rec, []models.Record(nil))
	rec = transactions.GetLast("Alice", 0)
	assert.NotEqual(t, rec, []models.Record(nil))
}

func TestTransaction_Save(t *testing.T) {

	transactions := NewTransactionService()
	id := transactions.Save("Alice", withdrawalCodeTest, 400, true)
	assert.Equal(t, 0, id)
	rec := transactions.GetLast("Alice", 1)

	assert.Equal(t, []models.Record{{
		ID:            id,
		OperationCode: withdrawalCodeTest,
		UserID:        "Alice",
		Amount:        400,
		Success:       true,
	},
	}, rec)

	id = transactions.Save("Alice", withdrawalCodeTest, 400, false)
	rec = transactions.GetLast("Alice", 0)

	assert.Equal(t, []models.Record{
		{
			ID:            id,
			OperationCode: withdrawalCodeTest,
			UserID:        "Alice",
			Amount:        400,
			Success:       false,
		},
		{
			ID:            0,
			OperationCode: withdrawalCodeTest,
			UserID:        "Alice",
			Amount:        400,
			Success:       true,
		}},

		rec)
}
