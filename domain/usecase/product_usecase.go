package usecase

import (
	"context"
	request2 "service-product/internal/http/create_prorduct/request"
	response2 "service-product/internal/http/create_prorduct/response"
	lRequest "service-product/internal/http/list_product/request"
	lResponse "service-product/internal/http/list_product/response"
)

type ProductUseCase interface {
	CreateProduct(ctx context.Context, request request2.CreateProductRequest) *response2.CreateProductResponse
	GetListProduct(ctx context.Context, request lRequest.ListProductRequest) *lResponse.ListProductsResponse
}
