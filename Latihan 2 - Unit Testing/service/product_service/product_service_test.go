package product_service

import (
	"net/http"
	"products/domain/product_domain"
	"products/utils/error_utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	tm            = time.Now()
	createProduct func(*product_domain.Product) (*product_domain.Product, error_utils.MessageErr)
	updateProduct func(*product_domain.Product) (*product_domain.Product, error_utils.MessageErr)
	getProducts   func() ([]*product_domain.Product, error_utils.MessageErr)
	deleteProduct func(int64) error_utils.MessageErr
)

type productDomainMock struct{}

func (p *productDomainMock) CreateProduct(productReq *product_domain.Product) (*product_domain.Product, error_utils.MessageErr) {
	return createProduct(productReq)
}

func (p *productDomainMock) UpdateProduct(productReq *product_domain.Product) (*product_domain.Product, error_utils.MessageErr) {
	return updateProduct(productReq)
}

func (p *productDomainMock) GetProducts() ([]*product_domain.Product, error_utils.MessageErr) {
	return getProducts()
}

func (p *productDomainMock) DeleteProduct(productId int64) error_utils.MessageErr {
	return deleteProduct(productId)
}

func (p *productDomainMock) Close() {}

func TestProductservice_CreateProduct_Success(t *testing.T) {
	product_domain.ProductDomain = &productDomainMock{}

	requestBody := &product_domain.Product{
		Name:  "Product Test",
		Price: 10.10,
		Stock: 5,
	}

	expectedVal := &product_domain.Product{
		Id:        1,
		Name:      "Product Test",
		Price:     10.10,
		Stock:     5,
		CreatedAt: tm,
	}

	createProduct = func(p *product_domain.Product) (*product_domain.Product, error_utils.MessageErr) {
		return expectedVal, nil
	}

	product, err := ProductService.CreateProduct(requestBody)

	assert.Nil(t, err)

	assert.NotNil(t, product)

	assert.EqualValues(t, expectedVal.Id, product.Id)
	assert.EqualValues(t, expectedVal.Name, product.Name)
	assert.EqualValues(t, expectedVal.Price, product.Price)
	assert.EqualValues(t, expectedVal.Stock, product.Stock)
}

func TestProductservice_CreateProduct_BadRequest(t *testing.T) {
	requestBody := &product_domain.Product{
		Price: 10.10,
		Stock: 5,
	}

	product, err := ProductService.CreateProduct(requestBody)

	assert.NotNil(t, err)
	assert.Nil(t, product)

	assert.EqualValues(t, 400, err.Status())
	assert.EqualValues(t, "name is required", err.Message())
	assert.EqualValues(t, "bad_request", err.Error())
}

func TestProductservice_CreateProduct_ServerError(t *testing.T) {
	product_domain.ProductDomain = &productDomainMock{}

	createProduct = func(p *product_domain.Product) (*product_domain.Product, error_utils.MessageErr) {
		return nil, error_utils.NewInternalServerErrorr("something went wrong")
	}

	requestBody := &product_domain.Product{
		Name:  "Product Test",
		Price: 10.10,
		Stock: 5,
	}

	product, err := ProductService.CreateProduct(requestBody)

	assert.NotNil(t, err)
	assert.Nil(t, product)

	assert.EqualValues(t, 500, err.Status())
	assert.EqualValues(t, "something went wrong", err.Message())
	assert.EqualValues(t, "server_error", err.Error())
}

// Latihan Unit Testing
func TestProductService_GetProducts_Success(t *testing.T) {
	product_domain.ProductDomain = &productDomainMock{}

	product := &product_domain.Product{
		Id:        1,
		Name:      "Product Test",
		Price:     10.10,
		Stock:     5,
		CreatedAt: tm,
	}

	expectedVal := []*product_domain.Product{
		product,
		product,
	}

	getProducts = func() ([]*product_domain.Product, error_utils.MessageErr) {
		return expectedVal, nil
	}

	result, err := ProductService.GetProducts()
	assert.Nil(t, err)

	assert.NotNil(t, product)

	assert.EqualValues(t, result[0].Id, product.Id)
	assert.EqualValues(t, result[0].Name, product.Name)
	assert.EqualValues(t, result[0].Price, product.Price)
	assert.EqualValues(t, result[0].Stock, product.Stock)

	assert.EqualValues(t, result[1].Id, product.Id)
	assert.EqualValues(t, result[1].Name, product.Name)
	assert.EqualValues(t, result[1].Price, product.Price)
	assert.EqualValues(t, result[1].Stock, product.Stock)
}

func TestProductService_GetProducts_InternalServerError(t *testing.T) {
	product_domain.ProductDomain = &productDomainMock{}

	getProducts = func() ([]*product_domain.Product, error_utils.MessageErr) {
		return nil, error_utils.NewInternalServerErrorr("something went wrong")
	}

	result, err := ProductService.GetProducts()

	assert.Nil(t, result)
	assert.NotNil(t, err)

	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "something went wrong", err.Message())
	assert.EqualValues(t, "server_error", err.Error())
}
