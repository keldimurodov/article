package repo

import (
	pb "projects/article/post-service/genproto/post"
)

// PostStorageI ...
type ProductStorageI interface {
	Create(pb *pb.Post) (*pb.Post, error)
	Get(req *pb.GetPostRequest) (*pb.Post, error)
	GetAll(*pb.GetAllRequest) (*pb.GetAllResponse, error)
	Update(*pb.Post) (*pb.Post, error)
	Delete(*pb.GetPostRequest) (*pb.Post, error)
}
