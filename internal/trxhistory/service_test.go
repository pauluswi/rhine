package trxhistory

import (
	"context"
	"testing"

	"github.com/pauluswi/rhine/internal/entity"
	"github.com/pauluswi/rhine/pkg/log"
	"github.com/stretchr/testify/assert"
)

//var errCRUD = errors.New("error crud")

func Test_service_TrxHistoryCycle(t *testing.T) {
	logger, _ := log.NewForTest()
	s := NewService(&mockRepository{}, &mockKafkaRepository{}, &mockRedisRepository{}, logger)

	ctx := context.Background()

	//save
	trx, err := s.Save(ctx, entity.InputSave{ID: 111, TrxID: "2222", CustomerID: "62816777777", CD: "c", Status: "0", Amount: 10000})
	assert.Nil(t, err)
	assert.NotEmpty(t, trx.ID)

	// get
	res, err := s.Get(ctx, trx.ID)
	assert.Nil(t, err)
	assert.NotEmpty(t, res.CustomerID)

}

type mockRepository struct {
	items []entity.TrxHistory
}

type mockKafkaRepository struct {
	items []entity.TrxHistory
}

type mockElasticRepository struct {
	items []entity.TrxHistory
}

type mockRedisRepository struct {
	items []entity.TrxHistory
}

func (m mockRepository) Get(ctx context.Context, id int) (*entity.TrxHistory, error) {
	var trxhistory entity.TrxHistory
	trxhistory.ID = id
	trxhistory.CustomerID = "0816333333"
	return &trxhistory, nil
}

func (m mockRepository) GetAll(ctx context.Context) ([]entity.TrxHistory, error) {
	var trxhistory []entity.TrxHistory
	return trxhistory, nil
}

func (m mockRepository) Save(ctx context.Context, paytoken entity.TrxHistory) error {
	return nil
}

func (m mockRepository) SaveByKafka(ctx context.Context, trxhistory entity.TrxHistory) error {
	return nil
}

func (m mockRepository) Update(ctx context.Context, paytoken entity.TrxHistory) error {
	return nil
}

func (m mockKafkaRepository) ReadMessage(res chan<- []byte) {
}

func (m mockKafkaRepository) WriteMessage(data *entity.TrxHistory) error {
	return nil
}

func (r *mockElasticRepository) GetByElastic(param entity.InputGet) ([]entity.ElasticTrxHistory, error) {
	var trxhistory []entity.ElasticTrxHistory
	return trxhistory, nil
}

func (r *mockElasticRepository) StoreToElastic(data entity.ElasticTrxHistory) error {
	return nil
}

func (r *mockElasticRepository) Update(data entity.ElasticTrxHistory, id int) error {
	return nil
}

func (r *mockElasticRepository) Delete(id int) error {
	return nil
}

func (r *mockRedisRepository) GetByIdFromRedis(id int) (*entity.TrxHistory, error) {
	var trxhistory entity.TrxHistory
	return &trxhistory, nil
}
func (r *mockRedisRepository) StoreToRedis(data *entity.TrxHistory) error {
	return nil
}

func (r *mockRedisRepository) Update(data entity.TrxHistory) error {
	return nil
}

func (r *mockRedisRepository) Delete(data entity.TrxHistory) error {
	return nil
}
