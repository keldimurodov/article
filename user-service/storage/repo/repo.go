package repo

import (
	pbu "projects/article/user-service/genproto/user"
)

// UserStorageI ...
type UserStoragePostgresI interface {
	Create(*pbu.User) (*pbu.User, error)
	Get(user *pbu.GetUserRequest) (*pbu.User, error)
	GetAll(req *pbu.GetAllRequest) (*pbu.GetAllResponse, error)
	Delete(user *pbu.GetUserRequest) (*pbu.User, error)
	Update(user *pbu.User) (*pbu.User, error)
}
