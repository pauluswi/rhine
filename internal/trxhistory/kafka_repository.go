package trxhistory

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/pauluswi/rhine/internal/entity"
	//"github.com/pauluswi/rhine/pkg/dbcontext"

	"github.com/segmentio/kafka-go"
)

type KafkaRepository interface {
	WriteMessage(data *entity.TrxHistory) error
	ReadMessage(res chan<- []byte)
}

type kafkaRepository struct {
	conn    *kafka.Conn
	url     string
	topic   string
	timeout time.Duration
}

func newKafkaConnection(URL, topic string, timeout int) (*kafka.Conn, error) {
	kafkaConn, err := kafka.DialLeader(context.Background(), "tcp", URL, topic, 0)
	if err != nil {
		return nil, err
	}
	return kafkaConn, err
}

func NewKafkaConnection(URL, topic string, timeout int) (KafkaRepository, error) {
	repo := &kafkaRepository{
		topic:   topic,
		url:     URL,
		timeout: time.Duration(timeout) * time.Second,
	}

	conn, err := newKafkaConnection(URL, topic, timeout)
	if err != nil {
		return nil, err
	}
	repo.conn = conn

	return repo, nil
}

func (k kafkaRepository) WriteMessage(data *entity.TrxHistory) error {
	msgs, err := json.Marshal(data)
	if err != nil {
		return err
	}

	k.conn.SetWriteDeadline(time.Now().Add(k.timeout))

	if _, err = k.conn.WriteMessages(
		kafka.Message{Value: msgs},
	); err != nil {
		return err
	}
	return nil
}

func (k kafkaRepository) ReadMessage(res chan<- []byte) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{k.url},
		Topic:     k.topic,
		Partition: 0,
		MinBytes:  10,
		MaxBytes:  10e3,
	})

	ctx := context.Background()
	lastOffset, _ := k.conn.ReadLastOffset() // get latest offset
	r.SetOffset(lastOffset)                  // set latest offset

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			log.Println("kafka-repo ReadMessage", err.Error())
			break
		}
		//fmt.Printf("message at offset %d: %s = %s at %v\n", m.Offset, string(m.Key), string(m.Value), m.Time)
		//fmt.Println("Kafka Message:", string(m.Value))
		res <- m.Value
	}

}
