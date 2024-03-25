package http

import (
	"log"
	"net/http"
	"service-product/helper"
	"service-product/internal/http/list_product/request"
)

func (a ProductHandler) GetListPlayer(w http.ResponseWriter, r *http.Request) {
	var (
		req request.ListProductRequest
	)
	req.SearchValue = r.FormValue("search_value")
	req.SearchType = r.FormValue("search_type")
	req.SortBy = r.FormValue("sort_by")
	req.SortDir = r.FormValue("sort_dir")
	log.Println(req)
	res := a.usecase.GetListProduct(r.Context(), req)
	helper.RespondWithJSON(w, http.StatusOK, res)
}
