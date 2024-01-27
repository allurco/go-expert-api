package database

import (
	"github.com/allurco/go-expert-api/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{
		DB: db,
	}
}

func (p *Product) CreateProduct(product *entity.Product) error {

	return p.DB.Create(product).Error
}

func (p *Product) GetProductById(id string) (*entity.Product, error) {
	var product entity.Product
	if err := p.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *Product) UpdateProduct(product *entity.Product) error {
	_, err := p.GetProductById(product.ID.String())
	if err != nil {
		return err
	}
	return p.DB.Save(product).Error
}

func (p *Product) DeleteProduct(id string) error {
	product, err := p.GetProductById(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(product).Error
}

func (p *Product) GetProducts(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	var err error
	if sort == "" && sort != "asc" && sort != "desc" {
		sort = "id asc"
	}

	if page != 0 && limit != 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("create_at " + sort).Find(&products).Error
	} else {
		err = p.DB.Order("created_at " + sort).Find(&products).Error
	}

	return products, err
}
