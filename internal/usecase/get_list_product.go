package usecase

import (
	"context"
	lRequest "service-product/internal/http/list_product/request"
	lResponse "service-product/internal/http/list_product/response"
)

func (a ProductUseCase) GetListProduct(ctx context.Context, request lRequest.ListProductRequest) *lResponse.ListProductsResponse {
	resProduct, err := a.ProductRepository.GetProductPlayer(ctx, request)
	if err != nil {
		return &lResponse.ListProductsResponse{
			Status: lResponse.Status{
				Code:    400,
				Message: err.Error(),
			},
		}
	}
	var playerData []lResponse.Products
	for _, product := range *resProduct {
		playerData = append(playerData, lResponse.Products{
			ID:        product.ID,
			Name:      product.Name,
			Type:      product.Type,
			Price:     product.Price,
			UpdatedAt: product.UpdatedAt,
		})
	}

	return &lResponse.ListProductsResponse{
		Status: lResponse.Status{
			Code:    200,
			Message: "Success",
		},

		Data: playerData,
	}
}
