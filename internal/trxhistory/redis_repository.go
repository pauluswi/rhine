package trxhistory

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/pauluswi/rhine/internal/entity"
)

type CacheRepository interface {
	GetByIdFromRedis(id int) (*entity.TrxHistory, error)
	StoreToRedis(data *entity.TrxHistory) error
}

type trxHistoryRedisRepository struct {
	client     *redis.Client
	expiration time.Duration
}

func newRedisClient(redisURL string) (*redis.Client, error) {
	// opt, err := redis.ParseURL("redis://:qwerty@localhost:6379/1")
	opt, e := redis.ParseURL(redisURL)
	if e != nil {
		return nil, e
	}
	client := redis.NewClient(opt)
	if _, e = client.Ping().Result(); e != nil {
		return nil, e
	}
	return client, e
}

func NewRedisRepository(redisURL string, expiration int) (CacheRepository, error) {
	repo := &trxHistoryRedisRepository{
		expiration: time.Duration(expiration) * time.Second,
	}
	client, e := newRedisClient(redisURL)
	if e != nil {
		return nil, e
	}
	repo.client = client
	return repo, nil
}

func (r *trxHistoryRedisRepository) GetByIdFromRedis(id int) (*entity.TrxHistory, error) {
	key := strconv.Itoa(id)
	data := entity.TrxHistory{}

	dataRedis, err := r.client.Get(key).Result()
	if err != nil {
		return &data, err
	}
	dataByte := []byte(dataRedis)
	if len(dataByte) == 0 {
		return &data, err
	}

	_res := new(entity.TrxHistory)

	if err := json.Unmarshal(dataByte, _res); err != nil {
		return &data, err
	}
	return _res, nil
}

func (r *trxHistoryRedisRepository) StoreToRedis(data *entity.TrxHistory) error {
	key := strconv.Itoa(data.ID)
	dataByte, err := data.MarshalBinary()
	if err != nil {
		return err
	}
	if _, err := r.client.Set(key, string(dataByte), 0).Result(); err != nil {
		return err
	}
	r.client.Expire(key, r.expiration)
	return nil
}
