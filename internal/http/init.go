package http

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"service-product/domain/usecase"
)

type ProductHandler struct {
	usecase usecase.ProductUseCase
}

func NewDeliveryHttpArea(r *mux.Router, usecase usecase.ProductUseCase) ProductHandler {
	handler := ProductHandler{
		usecase: usecase,
	}
	api := r.PathPrefix("/product").Subrouter()
	api.HandleFunc("", handler.GetListPlayer).Methods("GET")
	api.HandleFunc("/add", handler.CreateAccount).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
	return handler
}
