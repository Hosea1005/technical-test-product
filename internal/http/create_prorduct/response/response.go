package response

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type CreateProductResponse struct {
	Status Status `json:"status"`
}
