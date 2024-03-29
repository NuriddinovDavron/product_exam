package service

import (
	"context"
	pb "product_exam/genproto/product_exam"
	l "product_exam/pkg/logger"
	"product_exam/storage"

	grpcClient "product_exam/service/grpc_client"

	"github.com/jmoiron/sqlx"
)

// ProductService ...
type ProductService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpcClient.IServiceManager
}

func (p ProductService) CreateProduct(ctx context.Context, product *pb.CreateProductRequest) (*pb.Product, error) {
	respProduct, err := p.storage.Product().CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return respProduct, nil
}

func (p ProductService) GetProductById(ctx context.Context, request *pb.GetProductByIdRequest) (*pb.Product, error) {
	product, err := p.storage.Product().GetProductById(request)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p ProductService) GetAllProduct(ctx context.Context, request *pb.GetAllProductRequest) (*pb.GetAllProductResponse, error) {
	products, err := p.storage.Product().GetAllProduct(request)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p ProductService) UpdateProduct(ctx context.Context, product *pb.UpdateProductRequest) (*pb.Product, error) {
	respProduct, err := p.storage.Product().UpdateProduct(product)
	if err != nil {
		return nil, err
	}
	return respProduct, nil
}

func (p ProductService) DeleteProduct(ctx context.Context, request *pb.GetProductByIdRequest) (*pb.DeleteProductResponse, error) {
	err := p.storage.Product().DeleteProduct(request)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteProductResponse{
		Error: "successfully deleted",
	}, nil
}

func NewProductService(db *sqlx.DB, log l.Logger, client grpcClient.IServiceManager) *ProductService {
	return &ProductService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}
