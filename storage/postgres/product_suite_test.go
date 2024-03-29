package postgres

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"product_exam/config"
	pb "product_exam/genproto/product_exam"
	"product_exam/pkg/db"
	"product_exam/storage/repo"
	"testing"
)

type ProductRepositoryTestSuite struct {
	suite.Suite
	CleanUpFunc func()
	Repository  repo.ProductStorageI
}

func (s *ProductRepositoryTestSuite) SetupSuite() {
	pgPool, cleanUp := db.ConnectDBForSuite(config.Load())
	s.Repository = NewProductRepo(pgPool)
	s.CleanUpFunc = cleanUp
}

func (s *ProductRepositoryTestSuite) TestUserCRUD() {
	product := pb.CreateProductRequest{
		OwnerId:     uuid.NewString(),
		Name:        "product name",
		Description: "product description",
		Price:       45,
	}

	createdProduct, err := s.Repository.CreateProduct(&product)
	s.Suite.NotNil(createdProduct)
	s.Suite.NoError(err)
	s.Suite.Equal(product.Name, createdProduct.Name)
	s.Suite.Equal(product.Description, createdProduct.Description)

	getProductByIdReq := pb.GetProductByIdRequest{
		ProductId: createdProduct.Id,
	}

	getProduct, err := s.Repository.GetProductById(&getProductByIdReq)
	s.Suite.NotNil(getProduct)
	s.Suite.NoError(err)
	s.Suite.Equal(product.Name, getProduct.Name)
	s.Suite.Equal(product.Description, getProduct.Description)

	createdProduct.Name = "Updated Name"
	createdProduct.Description = "Updated Description"

	updatedProduct, err := s.Repository.GetProductById(&getProductByIdReq)
	s.Suite.NotNil(updatedProduct)
	s.Suite.NoError(err)

	getUpdatedProduct, err := s.Repository.GetProductById(&getProductByIdReq)
	s.Suite.NotNil(getUpdatedProduct)
	s.Suite.NoError(err)
	s.Suite.NotEqual(createdProduct.Name, getUpdatedProduct.Name)
	s.Suite.NotEqual(createdProduct.Description, getUpdatedProduct.Description)

	allProduct, err := s.Repository.GetAllProduct(&pb.GetAllProductRequest{Page: 1, Limit: 10})
	s.Suite.NotNil(allProduct)
	s.Suite.NoError(err)

	deleteProduct := s.Repository.DeleteProduct(&pb.GetProductByIdRequest{ProductId: createdProduct.Id})
	s.Suite.NotNil(deleteProduct)
	s.Suite.NoError(err)
}

func (s *ProductRepositoryTestSuite) TearDownSuite() {
	s.CleanUpFunc()
}

func TestProductRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ProductRepositoryTestSuite))
}
