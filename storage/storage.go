package storage

import (
	"product_exam/storage/postgres"
	"product_exam/storage/repo"

	"github.com/jmoiron/sqlx"
)

// IStorage ...
type IStorage interface {
	Product() repo.ProductStorageI
}

type Pg struct {
	db          *sqlx.DB
	productRepo repo.ProductStorageI
}

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *Pg {
	return &Pg{
		db:          db,
		productRepo: postgres.NewProductRepo(db),
	}
}

func (s Pg) Product() repo.ProductStorageI {
	return s.productRepo
}
