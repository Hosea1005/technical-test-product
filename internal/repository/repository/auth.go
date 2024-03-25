package repository

import (
	"go.uber.org/zap"
	"service-product/internal/config"
)

type ProductRepository struct {
	super  config.BaseRepository
	logger *zap.Logger
}

func NewProductRepository(super config.BaseRepository,
	logger *zap.Logger) *ProductRepository {
	return &ProductRepository{
		super:  super,
		logger: logger,
	}
}
func NewCacheRepository(super config.BaseRepository,
	logger *zap.Logger) *CacheRepository {
	return &CacheRepository{
		super:  super,
		logger: logger,
	}
}

type CacheRepository struct {
	super  config.BaseRepository
	logger *zap.Logger
}
