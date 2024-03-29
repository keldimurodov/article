package service

import (
	"context"
	"database/sql"
	u "projects/article/user-service/genproto/user"
	l "projects/article/user-service/pkg/logger"
	"projects/article/user-service/storage"

	"github.com/jmoiron/sqlx"
)

// UserService ...
type UserService struct {
	storage storage.IStorage
	logger  l.Logger
	db      *sql.DB
}

// NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *UserService) Create(ctx context.Context, req *u.User) (*u.User, error) {
	
	user, err := s.storage.User().Create(req)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Get(ctx context.Context, req *u.GetUserRequest) (*u.User, error) {
	user, err := s.storage.User().Get(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Update(ctx context.Context, req *u.User) (*u.User, error) {
	user, err := s.storage.User().Update(req)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Delete(ctx context.Context, req *u.GetUserRequest) (*u.User, error) {
	user, err := s.storage.User().Delete(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetAll(ctx context.Context, req *u.GetAllRequest) (*u.GetAllResponse, error) {
	users, err := s.storage.User().GetAll(req)
	if err != nil {
		return nil, err
	}
	return users, nil
}
