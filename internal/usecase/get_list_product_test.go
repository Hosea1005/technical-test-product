package usecase

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	domain "service-product/domain/repository"
	"service-product/helper"
	"service-product/internal/config"
	"service-product/internal/http/list_product/request"
	"service-product/internal/model"
	"testing"
)

func TestProductUseCase_GetListProduct(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		req := request.ListProductRequest{
			SearchValue: "",
			SearchType:  "",
			SortBy:      "",
			SortDir:     "",
		}
		prod := model.Product{
			ID:        helper.GeneratedUUID(),
			Name:      "asd",
			Type:      "asd",
			Price:     0,
			CreatedAt: nil,
			UpdatedAt: nil,
			DeletedAt: nil,
		}
		data := []model.Product{prod}
		productMock := new(domain.ProductMock)
		productMock.On("GetProductPlayer", mock.Anything, mock.Anything).Return(&data, nil)
		cacheMock := new(domain.CacheMock)
		var asd config.BaseRepository
		usecase := NewProductUseCase(nil, productMock, cacheMock, asd)

		_ = usecase.GetListProduct(context.Background(), req)
	})
	t.Run("Error - Insert Product", func(t *testing.T) {
		req := request.ListProductRequest{
			SearchValue: "",
			SearchType:  "",
			SortBy:      "",
			SortDir:     "",
		}
		productMock := new(domain.ProductMock)
		productMock.On("GetProductPlayer", mock.Anything, mock.Anything).Return(nil, errors.New("error to insert data"))
		cacheMock := new(domain.CacheMock)
		var asd config.BaseRepository
		usecase := NewProductUseCase(nil, productMock, cacheMock, asd)

		_ = usecase.GetListProduct(context.Background(), req)
	})
}
