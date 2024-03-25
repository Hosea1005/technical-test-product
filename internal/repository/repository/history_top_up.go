package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"log"
	"service-product/internal/http/create_prorduct/request"
)

func (r *CacheRepository) Set(ctx context.Context, key string, value request.CreateProductRequest) error {
	log.Println("Data : ", value)
	data, err := json.Marshal(value)
	if err != nil {
		//r.logger.Error("redis set key error marshal", zap.Error(err))
		return fmt.Errorf("failed to marshal value: %w", err)
	}

	if err := r.super.Redis.Set(ctx, key, data, 0).Err(); err != nil {
		r.logger.Error("redis set key error", zap.Error(err))
		return fmt.Errorf("failed to set key-value in redis: %w", err)
	}

	return nil
}
