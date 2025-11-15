package cmd

import (
	"ecommerce-server/handlers"
	"ecommerce-server/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager){
	mux.Handle("GET /test", manager.With(
		http.HandlerFunc(handlers.Test),
		middleware.Test2,
	))

	mux.Handle("GET /route", manager.With(
		http.HandlerFunc(handlers.Test),
		middleware.Test2,
	))

	mux.Handle("GET /products",manager.With(
		http.HandlerFunc(handlers.GetProducts),
		middleware.Test2,
	)) // route

	mux.Handle("POST /products",manager.With(
		http.HandlerFunc(handlers.CreateProduct),
		middleware.Test2,
	)) // route

	mux.Handle("GET /products/{productID}",manager.With(
		http.HandlerFunc(handlers.GetProductByID),
		middleware.Test2,
	)) // route
}