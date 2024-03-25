package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	lRequest "service-product/internal/http/list_product/request"
	"service-product/internal/model"
)

func (s *ProductRepository) GetProductPlayer(ctx context.Context, req lRequest.ListProductRequest) (*[]model.Product, error) {
	var products []model.Product
	db := s.super.PostgresSql

	// Memulai query dengan tabel produk
	query := db.Model(&model.Product{})

	// Pencarian berdasarkan SearchValue
	if req.SearchValue != "" {
		// Memeriksa apakah SearchValue adalah UUID yang valid
		if _, err := uuid.Parse(req.SearchValue); err == nil {
			// Jika valid, itu adalah pencarian berdasarkan ID
			query = query.Where("id = ?", req.SearchValue)
		} else {
			// Jika tidak valid, itu adalah pencarian berdasarkan nama
			query = query.Where("name LIKE ?", "%"+req.SearchValue+"%")
		}
	}

	// Pencarian berdasarkan SearchType
	if req.SearchType != "" {
		query = query.Where("type LIKE ?", "%"+req.SearchType+"%")

	}

	// Pengurutan berdasarkan SortBy dan SortDir
	if req.SortBy != "" {
		order := req.SortBy
		if req.SortDir != "" {
			if req.SortDir != "DESC" {
				order += " ASC"
			} else {
				order += " DESC"

			}

		}
		query = query.Order(order)
	}

	// Eksekusi query
	result := query.Find(&products)
	sqlString := query.Statement.SQL.String()
	fmt.Println("Query SQL yang dijalankan:", sqlString, result.Error)
	if result.Error != nil {
		return nil, result.Error
	}

	if len(products) == 0 {
		return nil, errors.New("products not found")
	}

	return &products, nil
}
