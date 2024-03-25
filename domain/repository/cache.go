package domain

import (
	"context"
	"service-product/internal/http/create_prorduct/request"
)

type CacheRepository interface {
	Set(ctx context.Context, key string, value request.CreateProductRequest) error
	SetToken(ctx context.Context, key string, token string) error
	GetToken(ctx context.Context, key string) (string, error)
	DeleteToken(ctx context.Context, key string) error
}
