package trxhistory

import (
	"context"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/pauluswi/rhine/internal/entity"
	"github.com/pauluswi/rhine/pkg/dbcontext"
	"github.com/pauluswi/rhine/pkg/log"
)

// Repository encapsulates the logic to access paytoken from the data source.
type Repository interface {
	// Get returns a transaction history information.
	Get(ctx context.Context, id int) (*entity.TrxHistory, error)
	// Save will store a transaction history information into data source.
	Save(ctx context.Context, trxhistory entity.TrxHistory) error
	// SaveByKafka will store a transaction history information into data source.
	SaveByKafka(ctx context.Context, trxhistory entity.TrxHistory) error
}

// repository persists paytoken in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new paytoken repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get returns a transaction history.
func (r repository) Get(ctx context.Context, id int) (*entity.TrxHistory, error) {
	var trxhistory entity.TrxHistory
	err := r.db.With(ctx).Select("id", "trx_id", "customer_id", "cd", "status", "amount", "created_at", "updated_at").
		From("trx_history").
		Where(dbx.HashExp{"id": id}).
		One(&trxhistory)
	return &trxhistory, err
}

// // Get returns a transaction history.
// func (r repository) GetAll(ctx context.Context) ([]entity.TrxHistory, error) {
// 	var trxhistory []entity.TrxHistory
// 	err := r.db.With(ctx).
// 		Select().
// 		All(&trxhistory)
// 	return trxhistory, err
// }

// Save will store a trxhistory information into data source.
func (r repository) Save(ctx context.Context, trxhistory entity.TrxHistory) error {
	_, err := r.db.With(ctx).Insert("trx_history", dbx.Params{
		"id":          trxhistory.ID,
		"trx_id":      trxhistory.TrxID,
		"customer_id": trxhistory.CustomerID,
		"cd":          trxhistory.CD,
		"status":      trxhistory.Status,
		"amount":      trxhistory.Amount,
		"created_at":  trxhistory.CreatedAt,
		"updated_at":  trxhistory.UpdatedAt,
	}).Execute()
	return err
}

// SaveByKafka will store a trxhistory information into data source.
func (r repository) SaveByKafka(ctx context.Context, trxhistory entity.TrxHistory) error {
	_, err := r.db.With(ctx).Insert("trx_history", dbx.Params{
		"id":          trxhistory.ID,
		"trx_id":      trxhistory.TrxID,
		"customer_id": trxhistory.CustomerID,
		"cd":          trxhistory.CD,
		"status":      trxhistory.Status,
		"amount":      trxhistory.Amount,
		"created_at":  trxhistory.CreatedAt,
		"updated_at":  trxhistory.UpdatedAt,
	}).Execute()
	return err
}
