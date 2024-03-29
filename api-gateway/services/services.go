package services

import (
	"fmt"

	"projects/article/api-gateway/config"
	p "projects/article/api-gateway/genproto/post"
	u "projects/article/api-gateway/genproto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	// r "projects/article/projects/article/api-gateway/genproto/recommend"
	// py "projects/article/projects/article/api-gateway/genproto/payments"
)

type IServiceManager interface {
	UserService() u.UserServiceClient
	PostService() p.PostServiceClient
	// RecommendService() r.RecommendServiceClient
	// PaymentService() py.PaymentServiceClient
}

type serviceManager struct {
	userService u.UserServiceClient
	postService p.PostServiceClient
	// recommendService r.RecommendServiceClient
	// paymentService py.PaymentServiceClient
}

func (s *serviceManager) UserService() u.UserServiceClient {
	return s.userService
}

func (s *serviceManager) PostService() p.PostServiceClient {
	return s.postService
}

// func (s *serviceManager) RecommendService() r.RecommendServiceClient {
// 	return s.recommendService
// }

// func (s *serviceManager) PaymentService() py.PaymentServiceClient {
// 	return s.paymentService
// }

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PostServiceHost, conf.PostServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	// connRecommend, err := grpc.Dial(
	// 	fmt.Sprintf("%s:%d", conf.RecommendServiceHost, conf.RecommendServicePort),
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	return nil, err
	// }

	// connPayment, err := grpc.Dial(
	// 	fmt.Sprintf("%s:%d", conf.PaymendServiceHost, conf.PaymendServicePort),
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	return nil, err
	// }

	serviceManager := &serviceManager{
		userService: u.NewUserServiceClient(connUser),
		postService: p.NewPostServiceClient(connPost),
	}

	return serviceManager, nil
}
