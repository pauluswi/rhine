package trxhistory

import (
	"context"
	"encoding/json"

	"github.com/pauluswi/rhine/internal/entity"
)

type KafkaService interface {
	ReadMessage(ctx context.Context, repo Repository, redisRepo CacheRepository) error
}

type kafkaService struct {
	repo      KafkaRepository
	redisRepo CacheRepository
}

func NewKafkaService(repo KafkaRepository, redisRepo CacheRepository) KafkaService {
	return &kafkaService{
		repo,
		redisRepo,
	}
}

func (u *kafkaService) ReadMessage(ctx context.Context, repo Repository, redisRepo CacheRepository) error {
	dataChan := make(chan []byte) // it will be sent to ReadMessage function

	go func() {
		for {
			select {
			case dataByte := <-dataChan:
				data := new(entity.TrxHistory)
				if err := json.Unmarshal(dataByte, data); err != nil {
					return
				}

				if err := repo.SaveByKafka(ctx, *data); err != nil {
					return
				}

				if err := redisRepo.StoreToRedis(data); err != nil {
					return
				}

			default:
			}
		}
	}()

	ctx.Done()
	u.repo.ReadMessage(dataChan)

	return nil
}
