package usecase

import (
	"context"
	"service-product/domain/entity"
	"service-product/helper"
	request2 "service-product/internal/http/create_prorduct/request"
	response2 "service-product/internal/http/create_prorduct/response"
	"time"
)

func (a ProductUseCase) CreateProduct(ctx context.Context, request request2.CreateProductRequest) *response2.CreateProductResponse {
	now := time.Now()
	payloadProduct, err := entity.NewProduct(&entity.ProductDTO{
		Id:          helper.GeneratedUUID(),
		Name:        request.Name,
		TypeProduct: request.Type,
		Price:       request.Price,
		CreatedAt:   &now,
		UpdatedAt:   &now,
	})
	if err != nil {
		return &response2.CreateProductResponse{
			Status: response2.Status{
				Code:    400,
				Message: err.Error(),
			},
		}
	}
	_, errInsert := a.ProductRepository.InsertProduct(ctx, payloadProduct)
	if errInsert != nil {
		return &response2.CreateProductResponse{
			Status: response2.Status{
				Code:    400,
				Message: errInsert.Error(),
			},
		}
	}
	errHistoryAddProduct := a.CacheService.Set(ctx, request.Name, request)
	if errHistoryAddProduct != nil {
		return &response2.CreateProductResponse{
			Status: response2.Status{
				Code:    400,
				Message: errHistoryAddProduct.Error(),
			},
		}
	}

	return &response2.CreateProductResponse{
		Status: response2.Status{
			Code:    200,
			Message: "Success Create Product",
		},
	}
}
