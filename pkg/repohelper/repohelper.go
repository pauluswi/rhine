package repohelper

import (
	"log"
	"strconv"

	"github.com/pauluswi/rhine/internal/config"

	repo "github.com/pauluswi/rhine/internal/trxhistory"
)

func KafkaConnection(cfg *config.Config) repo.KafkaRepository {
	timeout, _ := strconv.Atoi(cfg.KAFKA_TIMEOUT)
	if timeout == 0 {
		timeout = 10
	}
	url := cfg.KAFKA_URL
	if url == "" {
		url = "localhost:9092"
	}
	topic := cfg.KAFKA_TOPIC
	if topic == "" {
		topic = "trxhistory"
	}
	repo, err := repo.NewKafkaConnection(url, topic, timeout)
	if err != nil {
		log.Fatal(err)
	}
	return repo
}

func RedisConnection(cfg *config.Config) repo.CacheRepository {
	timeout, _ := strconv.Atoi(cfg.REDIS_EXPIRED)
	if timeout == 0 {
		timeout = 10
	}
	url := cfg.REDIS_URL
	if url == "" {
		url = "redis://:@localhost:6379/0"
	}
	repo, err := repo.NewRedisRepository(url, timeout)
	if err != nil {
		log.Fatal(err)
	}
	return repo
}
