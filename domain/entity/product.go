package entity

import (
	"github.com/google/uuid"
	"time"
)

type (
	ProductDTO struct {
		Id          uuid.UUID
		Name        string
		TypeProduct string
		Price       float64
		CreatedAt   *time.Time
		UpdatedAt   *time.Time
		DeletedAt   *time.Time
	}
	Product struct {
		id          uuid.UUID
		name        string
		typeProduct string
		price       float64
		createdAt   *time.Time
		updatedAt   *time.Time
		deletedAt   *time.Time
	}
)

func NewProduct(payload *ProductDTO) (*Product, error) {
	return RebuildProduct(payload), nil
}
func RebuildProduct(payload *ProductDTO) *Product {
	return &Product{
		id:          payload.Id,
		name:        payload.Name,
		typeProduct: payload.TypeProduct,
		price:       payload.Price,
		createdAt:   payload.CreatedAt,
		updatedAt:   payload.UpdatedAt,
		deletedAt:   payload.DeletedAt,
	}
}

func (p *Product) Id() uuid.UUID {
	return p.id
}
func (p *Product) Name() string {
	return p.name
}
func (p *Product) Type() string {
	return p.typeProduct
}
func (p *Product) Price() float64 {
	return p.price
}
func (p *Product) CreatedAt() *time.Time {
	return p.createdAt
}
func (p *Product) UpdatedAt() *time.Time {
	return p.updatedAt
}
func (p *Product) DeletedAt() *time.Time {
	return p.deletedAt
}
