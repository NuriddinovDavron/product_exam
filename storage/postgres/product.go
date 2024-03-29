package postgres

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	pb "product_exam/genproto/product_exam"
)

type ProductRepo struct {
	db *sqlx.DB
}

func (p ProductRepo) CreateProduct(mailReq *pb.CreateProductRequest) (*pb.Product, error) {

	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	productId := id.String()
	var res pb.Product
	query := `INSERT INTO products_exam (id, owner_id, name, description, price) VALUES ($1, $2, $3, $4, $5) RETURNING id, owner_id, name, description, price`
	err = p.db.QueryRow(query, productId, mailReq.OwnerId, mailReq.Name, mailReq.Description, mailReq.Price).Scan(
		&res.Id,
		&res.OwnerId,
		&res.Name,
		&res.Description,
		&res.Price)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (p ProductRepo) GetProductById(isUnReq *pb.GetProductByIdRequest) (*pb.Product, error) {
	var res pb.Product
	query := `SELECT id, owner_id, name, description, price, created_at, updated_at FROM products_exam WHERE id=$1`
	err := p.db.QueryRow(query, isUnReq.ProductId).Scan(
		&res.Id,
		&res.OwnerId,
		&res.Name,
		&res.Description,
		&res.Price,
		&res.CreatedAt,
		&res.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (p ProductRepo) GetAllProduct(crUsReq *pb.GetAllProductRequest) (*pb.GetAllProductResponse, error) {
	var allUser pb.GetAllProductResponse
	query := `select id, owner_id, name, description, price from products_exam limit $1 offset $2`
	offset := crUsReq.Limit * (crUsReq.Page - 1)
	rows, err := p.db.Query(query, crUsReq.Limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user pb.Product
		err := rows.Scan(&user.Id, &user.OwnerId, &user.Name, &user.Description, &user.Price)
		if err != nil {
			return nil, err
		}
		allUser.Products = append(allUser.Products, &user)
	}
	return &allUser, nil
}

func (p ProductRepo) UpdateProduct(logInReq *pb.UpdateProductRequest) (*pb.Product, error) {
	query := `update products_exam set owner_id=$1, name=$2, description=$3, price=$4 where id=$5 returning id, owner_id, name, description, price`
	var res pb.Product
	err := p.db.QueryRow(query, logInReq.OwnerId, logInReq.Name, logInReq.Description, logInReq.Price, logInReq.Id).Scan(
		&res.Id,
		&res.OwnerId,
		&res.Name,
		&res.Description,
		&res.Price)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (p ProductRepo) DeleteProduct(logInReq *pb.GetProductByIdRequest) error {
	query := `delete from products_exam where id=$1`
	err := p.db.QueryRow(query, logInReq.ProductId).Err()
	if err != nil {
		return err
	}
	return nil
}

func NewProductRepo(db *sqlx.DB) *ProductRepo {
	return &ProductRepo{db: db}
}
