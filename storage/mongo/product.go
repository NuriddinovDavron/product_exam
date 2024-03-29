package mongo

import (
	"context"
	pb "product_exam/genproto/product_exam"
	"product_exam/pkg/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productRepo struct {
	database *mongo.Database
	log      logger.Logger
}

func NewProductRepo(database *mongo.Database, log logger.Logger) *productRepo {
	return &productRepo{database: database, log: log}
}

func (p *productRepo) Create(ctx context.Context, req *pb.CreateProductRequest) (*pb.Product, error) {
	collection := p.database.Collection("products")
	result, err := collection.InsertOne(ctx, req)
	if err != nil {
		return nil, err
	}

	var response pb.Product
	filter := bson.M{"_id": result.InsertedID}
	err = collection.FindOne(ctx, filter).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (p *productRepo) GetProductById(ctx context.Context, req *pb.GetProductByIdRequest) (*pb.Product, error) {

	collection := p.database.Collection("products")

	var response pb.Product
	filter := bson.M{"id": req.ProductId}
	err := collection.FindOne(ctx, filter).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (p *productRepo) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.Product, error) {
	collection := p.database.Collection("products")

	var response pb.Product

	filter := bson.M{"id": req.Id}

	updateReq := bson.M{
		"$set": bson.M{
			"name":        req.Name,
			"description": req.Description,
			"price":       req.Price,
			"updated_at":  time.Now(),
		},
	}

	err := collection.FindOneAndUpdate(ctx, filter, updateReq).Decode(&req)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (p *productRepo) DeleteProduct(ctx context.Context, req *pb.GetProductByIdRequest) error {
	collection := p.database.Collection("products")

	filter := bson.M{"id": req.ProductId}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return  err
	}

	return nil
}

func (p *productRepo) GetAllProduct(ctx context.Context, req *pb.GetAllProductRequest) (*pb.GetAllProductResponse, error) {
	collection := p.database.Collection("products")

	var response pb.GetAllProductResponse

	reqOptions := options.Find()

	reqOptions.SetSkip(int64(req.Page-1) * int64(req.Limit))
	reqOptions.SetLimit(int64(req.Limit))

	cursor, err := collection.Find(ctx, bson.M{}, reqOptions)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var product pb.Product
		err = cursor.Decode(&product)
		if err != nil {
			return nil, err
		}

		response.Products = append(response.Products, &product)
	}

	return &response, nil
}
