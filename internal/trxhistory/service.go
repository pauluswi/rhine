package trxhistory

import (
	"context"
	"fmt"
	"time"

	"github.com/pauluswi/rhine/internal/entity"
	"github.com/pauluswi/rhine/pkg/log"
)

// Service encapsulates usecase logic for albums.
type Service interface {
	Get(ctx context.Context, id int) (*entity.TrxHistory, error)
	Save(ctx context.Context, req entity.InputSave) (out entity.OutSave, err error)
}

// TrxHistory represents the data about a transaction history.
type TrxHistory struct {
	entity.TrxHistory
}

type service struct {
	repo      Repository
	kafkaRepo KafkaRepository
	redisRepo CacheRepository
	logger    log.Logger
}

// NewService creates a new payment transaction history.
func NewService(repo Repository, kafkaRepo KafkaRepository, redisRepo CacheRepository, logger log.Logger) Service {
	return service{repo, kafkaRepo, redisRepo, logger}
}

// --- list of error and constants
var (
	ErrValidation         = fmt.Errorf("validation error")
	ErrDBPersist          = fmt.Errorf("persist to database error")
	ErrTrxHistoryNotFound = fmt.Errorf("transaction history not found in database")
)

// Get a transaction history from ID param
func (s service) Get(ctx context.Context, id int) (out *entity.TrxHistory, err error) {
	outData := entity.TrxHistory{}
	data, err := s.redisRepo.GetByIdFromRedis(id)
	// if no data on cache then get from DB
	if err != nil || data == &outData {
		trxhistory, err := s.repo.Get(ctx, id)
		if err != nil {
			return &outData, err
		}
		return trxhistory, nil
	}

	return data, nil
}

// Create a transaction history
func (s service) Save(ctx context.Context, req entity.InputSave) (out entity.OutSave, err error) {
	defer func() {
		if err != nil {
			s.logger.Error(ctx, err.Error())
		}
	}()

	now := time.Now().UTC()

	// build trxhistory
	trxhistory := entity.NewTrxHistory()
	trxhistory.ID = req.ID
	trxhistory.TrxID = req.TrxID
	trxhistory.CustomerID = req.CustomerID
	trxhistory.CD = req.CD
	trxhistory.Status = req.Status
	trxhistory.Amount = req.Amount
	trxhistory.CreatedAt = now
	trxhistory.UpdatedAt = now

	// Write Kafka message
	err = s.kafkaRepo.WriteMessage(trxhistory)
	if err != nil {
		return entity.OutSave{}, err
	}

	output := entity.OutSave{
		ID: trxhistory.ID,
	}
	return output, err
}
