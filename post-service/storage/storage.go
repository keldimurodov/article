package storage

import (
	"projects/article/post-service/storage/postgres"
	"projects/article/post-service/storage/repo"

	"github.com/jmoiron/sqlx"
)

// IStorage ...
type IStorage interface {
	Product() repo.ProductStorageI
}

type storagePg struct {
	db          *sqlx.DB
	productRepo repo.ProductStorageI
}

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:          db,
		productRepo: postgres.NewPostRepo(db),
	}
}

func (s storagePg) Product() repo.ProductStorageI {
	return s.productRepo
}
