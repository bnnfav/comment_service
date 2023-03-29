package grpcclient

import (
	"bay_store/comment_service/config"
	co "bay_store/comment_service/genproto/order"
	cu "bay_store/comment_service/genproto/user"
	cp "bay_store/comment_service/genproto/product"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients interface {
	User() cu.UserServiceClient
	Order() co.OrderServiceClient
	Product() cp.ProductServiceClient
}

type ServiceManager struct {
	Config         config.Config
	userService    cu.UserServiceClient
	orderService   co.OrderServiceClient
	productService cp.ProductServiceClient
}

func New(cfg config.Config) (*ServiceManager, error) {
	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("user service dial host:%s, port:%s", cfg.UserServiceHost, cfg.UserServicePort)
	}

	connOrder, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.OrderServiceHost, cfg.OrderServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("order service dial host:%s, port:%s", cfg.OrderServiceHost, cfg.OrderServicePort)
	}

	connProduct, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.ProductServiceHost, cfg.ProductServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("product service dial host:%s, port:%s", cfg.ProductServiceHost, cfg.ProductServicePort)
	}

	return &ServiceManager{
		Config:         cfg,
		userService:    cu.NewUserServiceClient(connUser),
		orderService:   co.NewOrderServiceClient(connOrder),
		productService: cp.NewProductServiceClient(connProduct),
	}, nil
}

func (s *ServiceManager) User() cu.UserServiceClient {
	return s.userService
}

func (s *ServiceManager) Order() co.OrderServiceClient {
	return s.orderService
}

func (s *ServiceManager) Product() cp.ProductServiceClient {
	return s.productService
}
