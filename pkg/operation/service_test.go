package operation

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"facade/pkg/models"
)

const (
	withdrawalCodeTest = iota
	balanceCodeTest
)

func TestTransaction_Get(t *testing.T) {

	transactions := NewService()
	transactions.Save("Alice", withdrawalCodeTest, 400, true)
	rec := transactions.GetLast("Bob", 0)
	assert.Equal(t, rec, []models.Record(nil))
	rec = transactions.GetLast("Alice", 0)
	assert.NotEqual(t, rec, []models.Record(nil))
}

func TestTransaction_Save(t *testing.T) {

	transactions := NewService()
	id := transactions.Save("Alice", balanceCodeTest, 400, true)
	assert.Equal(t, 0, id)
	rec := transactions.GetLast("Alice", 1)

	assert.Equal(t, []models.Record{{
		ID:            id,
		OperationCode: balanceCodeTest,
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
			OperationCode: balanceCodeTest,
			UserID:        "Alice",
			Amount:        400,
			Success:       true,
		}},

		rec)
}
