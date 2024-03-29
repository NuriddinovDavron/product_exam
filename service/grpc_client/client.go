package grpc_client

import (
	"fmt"
	"google.golang.org/grpc"
	"product_exam/config"
	pbp "product_exam/genproto/user_exam"
)

type IServiceManager interface {
	UserService() pbp.UserServiceClient
}

type serviceManager struct {
	cfg         config.Config
	userService pbp.UserServiceClient
}

func (s *serviceManager) UserService() pbp.UserServiceClient {
	return s.userService
}

func New(cfg config.Config) (IServiceManager, error) {
	// dail to post-service
	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("product service dail host: %s port : %d", cfg.UserServiceHost, cfg.UserServicePort)
	}

	return &serviceManager{
		cfg:         cfg,
		userService: pbp.NewUserServiceClient(connUser),
	}, nil
}
