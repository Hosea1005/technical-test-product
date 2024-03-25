package repository

import (
	"context"
	"service-product/domain/entity"
	"service-product/internal/model"
)

func (s *ProductRepository) InsertProduct(ctx context.Context, data *entity.Product) (*entity.Product, error) {
	account := model.Product{
		ID:        data.Id(),
		Name:      data.Name(),
		Type:      data.Type(),
		Price:     data.Price(),
		CreatedAt: data.CreatedAt(),
		UpdatedAt: data.UpdatedAt(),
	}
	result := s.super.PostgresSql.Create(&account)
	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}
