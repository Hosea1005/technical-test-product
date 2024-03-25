package main

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	dUsecase "service-product/domain/usecase"
	"service-product/internal/config"
	httpHandler "service-product/internal/http"
	"service-product/internal/repository/repository"
	"service-product/internal/usecase"
)

func main() {
	logger := config.Logger()
	config.Environment()
	logger.Info("initializing service product ")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	repo := config.BaseRepository{
		PostgresSql: config.PostgresDB(logger),
		Redis:       config.Redis(ctx, logger),
	}
	// Initialize repository instances
	userRepo := repository.NewProductRepository(repo, logger)
	cacheRepo := repository.NewCacheRepository(repo, logger)
	// Initialize use case instances
	userService := usecase.NewProductUseCase(logger, userRepo, cacheRepo, repo)

	r := mux.NewRouter()
	initHandler(r, userService)
	http.Handle("/", r)

}
func initHandler(r *mux.Router, usecase dUsecase.ProductUseCase) {
	httpHandler.NewDeliveryHttpArea(r, usecase)
}
