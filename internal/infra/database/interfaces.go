package database

import "github.com/allurco/go-expert-api/internal/entity"

type UserInterface interface {
	CreateUser(user *entity.User) error
	GetUserByEmail(email string) (*entity.User, error)
}

type ProductInterface interface {
	CreateProduct(product *entity.Product) error
	GetProductById(id string) (*entity.Product, error)
	GetProducts(page, limit int, sort string) ([]*entity.Product, error)
	UpdateProduct(product *entity.Product) error
	DeleteProduct(id string) error
}
