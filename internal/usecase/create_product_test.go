package usecase

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	domain "service-product/domain/repository"
	"service-product/internal/config"
	"service-product/internal/http/create_prorduct/request"
	"testing"
)

func TestProductUseCase_CreateProduct(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		req := request.CreateProductRequest{
			Name:  "",
			Type:  "",
			Price: 0,
		}
		productMock := new(domain.ProductMock)
		productMock.On("InsertProduct", mock.Anything, mock.Anything).Return(nil, nil)
		cacheMock := new(domain.CacheMock)
		cacheMock.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		var asd config.BaseRepository
		usecase := NewProductUseCase(nil, productMock, cacheMock, asd)

		_ = usecase.CreateProduct(context.Background(), req)
	})
	t.Run("Error - Insert Product", func(t *testing.T) {
		req := request.CreateProductRequest{
			Name:  "",
			Type:  "",
			Price: 0,
		}
		productMock := new(domain.ProductMock)
		productMock.On("InsertProduct", mock.Anything, mock.Anything).Return(nil, errors.New("error to insert data"))
		cacheMock := new(domain.CacheMock)
		cacheMock.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		var asd config.BaseRepository
		usecase := NewProductUseCase(nil, productMock, cacheMock, asd)

		_ = usecase.CreateProduct(context.Background(), req)
	})
	t.Run("Error - setProduct", func(t *testing.T) {
		req := request.CreateProductRequest{
			Name:  "",
			Type:  "",
			Price: 0,
		}
		productMock := new(domain.ProductMock)
		productMock.On("InsertProduct", mock.Anything, mock.Anything).Return(nil, nil)
		cacheMock := new(domain.CacheMock)
		cacheMock.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error to insert data"))
		var asd config.BaseRepository
		usecase := NewProductUseCase(nil, productMock, cacheMock, asd)

		_ = usecase.CreateProduct(context.Background(), req)
	})
}
