package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/allurco/go-expert-api/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	db, err := CreateConnection(&entity.Product{})
	if err != nil {
		t.Error(err)
	}
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

func TestUpdateProduct(t *testing.T) {
	db, err := CreateConnection(&entity.Product{})
	if err != nil {
		t.Error(err)
	}
	product, _ := entity.NewProduct("Product 1", 10)
	db.Create(product)

	product.Name = "Product 2"

	productDB := NewProduct(db)
	err = productDB.UpdateProduct(product)
	assert.Nil(t, err)

	var newProduct *entity.Product
	newProduct, err = productDB.GetProductById(product.ID.String())
	assert.Nil(t, err)

	assert.Equal(t, "Product 2", newProduct.Name)

}

func TestDeleteProduct(t *testing.T) {
	db, err := CreateConnection(&entity.Product{})
	if err != nil {
		t.Error(err)
	}
	product, _ := entity.NewProduct("Product 1", 10)

	productDB := NewProduct(db)
	err = productDB.CreateProduct(product)
	assert.Nil(t, err)

	var savedProduct *entity.Product
	savedProduct, err = productDB.GetProductById(product.ID.String())
	assert.Nil(t, err)

	assert.Equal(t, product.ID.String(), savedProduct.ID.String())

	err = productDB.DeleteProduct(product.ID.String())
	assert.Nil(t, err)

	_, err = productDB.GetProductById(product.ID.String())
	assert.NotNil(t, err)

}

func TestGetProductById(t *testing.T) {
	db, err := CreateConnection(&entity.Product{})
	if err != nil {
		t.Error(err)
	}
	product, _ := entity.NewProduct("Product 1", 10)

	productDB := NewProduct(db)
	err = productDB.CreateProduct(product)
	assert.Nil(t, err)

	var foundProduct *entity.Product
	foundProduct, err = productDB.GetProductById(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, product.Name, foundProduct.Name)

}

func TestGetProducts(t *testing.T) {
	db, err := CreateConnection(&entity.Product{})
	if err != nil {
		t.Error(err)
	}
	productDB := NewProduct(db)
	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product Number %d", i), rand.Float64()*100)
		assert.Nil(t, err)
		db.Create(product)
	}

	var products []entity.Product
	assert.Len(t, products, 0)

	products, err = productDB.GetProducts(1, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product Number 1", products[0].Name)
	assert.Equal(t, "Product Number 10", products[9].Name)

	products, err = productDB.GetProducts(2, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product Number 11", products[0].Name)
	assert.Equal(t, "Product Number 20", products[9].Name)

	products, err = productDB.GetProducts(3, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product Number 21", products[0].Name)
	assert.Equal(t, "Product Number 23", products[2].Name)

}
