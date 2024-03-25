package domain

import (
	"context"
	"github.com/google/uuid"
	"service-product/domain/entity"
	lRequest "service-product/internal/http/list_product/request"
	"service-product/internal/model"
)

//go:generate mockery --name ProductRepository --filename product_mock.go --inpackage --structname ProductMock --output=./
type ProductRepository interface {
	GetProductPlayer(ctx context.Context, req lRequest.ListProductRequest) (*[]model.Product, error)
	CheckIdProduct(ctx context.Context, id uuid.UUID) error
	InsertProduct(ctx context.Context, data *entity.Product) (*entity.Product, error)
}
