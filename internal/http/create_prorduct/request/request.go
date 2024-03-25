package request

type CreateProductRequest struct {
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}
