package repo

import (
	pb "product_exam/genproto/product_exam"
)

// ProductStorageI ...
type ProductStorageI interface {
	CreateProduct(mailReq *pb.CreateProductRequest) (*pb.Product, error)
	GetProductById(isUnReq *pb.GetProductByIdRequest) (*pb.Product, error)
	GetAllProduct(crUsReq *pb.GetAllProductRequest) (*pb.GetAllProductResponse, error)
	UpdateProduct(logInReq *pb.UpdateProductRequest) (*pb.Product, error)
	DeleteProduct(logInReq *pb.GetProductByIdRequest) error
}
