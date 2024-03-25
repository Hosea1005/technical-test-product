package request

type ListProductRequest struct {
	SearchValue string
	SearchType  string
	SortBy      string
	SortDir     string
}
