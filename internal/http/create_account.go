package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"service-product/helper"
	"service-product/internal/http/create_prorduct/request"
)

func (a ProductHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var (
		req request.CreateProductRequest
	)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		helper.RespondWithJSON(w, http.StatusBadRequest, []error{errors.New("unauthorized")})
		return
	}
	defer r.Body.Close()
	res := a.usecase.CreateProduct(r.Context(), req)
	helper.RespondWithJSON(w, http.StatusOK, res)
}
