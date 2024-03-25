package usecase

import (
	"go.uber.org/zap"
	domain "service-product/domain/repository"
	"service-product/domain/usecase"
	"service-product/internal/config"
)

type ProductUseCase struct {
	Logger            *zap.Logger
	ProductRepository domain.ProductRepository
	CacheService      domain.CacheRepository
	Super             config.BaseRepository
}

func NewProductUseCase(
	Logger *zap.Logger,
	ProductUserRepository domain.ProductRepository,
	CacheService domain.CacheRepository,
	Super config.BaseRepository,
) usecase.ProductUseCase {
	return &ProductUseCase{Logger: Logger, ProductRepository: ProductUserRepository, CacheService: CacheService, Super: Super}
}
