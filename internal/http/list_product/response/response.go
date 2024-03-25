package response

import (
	"github.com/google/uuid"
	"time"
)

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type ListProductsResponse struct {
	Status Status     `json:"status"`
	Data   []Products `json:"data"`
}

type Products struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Type      string     `json:"type"`
	Price     float64    `json:"price"`
	UpdatedAt *time.Time `json:"updated_at_at"`
}
