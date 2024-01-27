package database

import (
	"testing"

	"github.com/allurco/go-expert-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	product, _ := entity.NewProduct("Product 1", 10)

	productDB := NewProduct(db)
	err = productDB.CreateProduct(product)
	assert.Nil(t, err)

	var productFound entity.Product
	err = db.First(&productFound, "id = ?", product.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)

}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	product, _ := entity.NewProduct("Product 1", 10)

	productDB := NewProduct(db)
	err = productDB.CreateProduct(product)
	assert.Nil(t, err)

	err = productDB.DeleteProduct(product.ID.String())
	assert.Nil(t, err)

}

func TestGetProductById(t *testing.T) {

}
