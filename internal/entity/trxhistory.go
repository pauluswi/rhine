package entity

import (
	"encoding/json"
	"time"
)

// TrxHistory struct is defined here
type TrxHistory struct {
	ID         int       `db:"id" validate:"required" bson:"_id" msgpack:"_id"`
	TrxID      string    `db:"trx_id" validate:"required" bson:"trx" msgpack:"trx"`
	CustomerID string    `db:"customer_id" validate:"required" bson:"customer" msgpack:"customer"`
	CD         string    `db:"cd" validate:"required" bson:"cd" msgpack:"cd"`
	Status     string    `db:"status" validate:"required" bson:"status" msgpack:"status"`
	Amount     int32     `db:"amount" validate:"required" bson:"amount" msgpack:"amount"`
	CreatedAt  time.Time `db:"created_at" validate:"required" bson:"created" msgpack:"created"`
	UpdatedAt  time.Time `db:"updated_at" validate:"required" bson:"updated" msgpack:"updated"`
}

func NewTrxHistory() *TrxHistory {
	return &TrxHistory{}
}

func (m *TrxHistory) TableName() string {
	return "trx_history"
}

type ElasticTrxHistory struct {
	ID      int       `json:"id" bson:"id" msgpack:"id"`
	Created time.Time `json:"created" bson:"created" msgpack:"created"`
}

func (m *ElasticTrxHistory) TableName() string {
	return "trx_history"
}

func (m TrxHistory) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}

func (m *TrxHistory) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

// --- I/O for Service function

// InputGet .
type InputGet struct {
	Filter map[string]interface{} `json:"filter" bson:"filter" msgpack:"filter"`
	Offset int                    `json:"offset" bson:"offset" msgpack:"offset"`
	Limit  int                    `json:"limit" bson:"limit" msgpack:"limit"`
	Order  map[string]bool        `json:"order" bson:"order" msgpack:"order"`
}

// InputSave .
type InputSave struct {
	ID         int    `json:"id" validate:"required"`
	TrxID      string `json:"trx_id" validate:"required,numeric"`
	CustomerID string `json:"customer_id" validate:"required,numeric,startswith=62,min=10"`
	CD         string `json:"cd" validate:"required,min=1"`
	Status     string `json:"status" validate:"required"`
	Amount     int32  `json:"amount" validate:"required,numeric"`
}

// OutSave .
type OutSave struct {
	ID int `json:"id"`
}
