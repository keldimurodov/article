package storage

import (
	"projects/article/user-service/storage/postgres"
	"projects/article/user-service/storage/repo"

	"github.com/jmoiron/sqlx"
)

// IStorage ...
type IStorage interface {
	User() repo.UserStoragePostgresI
}


type storagePg struct {
	db       *sqlx.DB
	userRepo repo.UserStoragePostgresI
}

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		userRepo: postgres.NewUserRepo(db),
	}
}

func (s storagePg) User() repo.UserStoragePostgresI {
	return s.userRepo
}
