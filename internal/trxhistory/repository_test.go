package trxhistory

import (
	"context"
	"testing"
	"time"

	"github.com/pauluswi/rhine/internal/entity"
	"github.com/pauluswi/rhine/internal/test"
	"github.com/pauluswi/rhine/pkg/log"
	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	logger, _ := log.NewForTest()
	db := test.DB(t)
	test.ResetTables(t, db, "trx_history")
	repo := NewRepository(db, logger)

	ctx := context.Background()

	//id, _ := uuid.NewV4()

	var trxData entity.TrxHistory
	trxData.ID = 1
	trxData.TrxID = "1111111"
	trxData.CustomerID = "6281100099"
	trxData.CD = "c"
	trxData.Status = "0"
	trxData.Amount = 25000
	trxData.CreatedAt = time.Now().UTC()
	trxData.UpdatedAt = time.Now().UTC()

	// create
	err := repo.Save(ctx, trxData)
	assert.Nil(t, err)

	// get
	trx, err := repo.Get(ctx, trxData.ID)
	assert.Nil(t, err)
	assert.Equal(t, "6281100099", trx.CustomerID)

}
