package test

import (
	"testing"

	"github.com/donluidi/goapi/internal/entity"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product := entity.NewProduct("Product", 10)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.GetId())
	assert.Equal(t, "Product", product.GetName())
	assert.Equal(t, 10, product.GetPrice())
}
