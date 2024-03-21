package repo

import (
	u "projects/article/user-service/genproto/user"
)

// UserStorageI ...
type UserStorageMongoI interface {
	Create(*u.User) (*u.User, error)
	Get(user *u.GetUserRequest) (*u.User, error)
	GetAll(req *u.GetAllRequest) (*u.GetAllResponse, error)
	Delete(user *u.GetUserRequest) (*u.User, error)
	Update(user *u.User) (*u.User, error)
}
