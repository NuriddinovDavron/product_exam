package mongorepo

import (
	pb "product_exam/genproto/product_exam"
	"context"
)

// ProductStorageI ...
type ProductStorageI interface {
	CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.Product, error)
	GetProductById(ctx context.Context, req *pb.GetProductByIdRequest) (*pb.Product, error)
	GetAllProduct(ctx context.Context, req *pb.GetAllProductRequest) (*pb.GetAllProductResponse, error)
	UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.Product, error)
	DeleteProduct(ctx context.Context, req *pb.GetProductByIdRequest) error
}