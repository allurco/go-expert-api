package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Product 1", 10.5)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.CreatedAt)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, 10.5, product.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 10.5)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("Product 1", 0)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	product, err := NewProduct("Product 1", -10.5)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrorInvalidPrice, err)
}

func TestProductValidate(t *testing.T) {
	product, err := NewProduct("Product 1", 10.5)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Nil(t, product.Validate())
}
