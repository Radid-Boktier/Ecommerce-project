package cmd

import (
	"ecommerce-server/global_router"
	"ecommerce-server/handlers"
	"ecommerce-server/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()

	mux := http.NewServeMux() // router

	mux.Handle("GET /test", manager.With(
		http.HandlerFunc(handlers.Test),
		middleware.Logger,
	))

	mux.Handle("GET /route", manager.With(
		http.HandlerFunc(handlers.Test),
		middleware.Logger,
	))

	mux.Handle("GET /products",manager.With(
		http.HandlerFunc(handlers.GetProducts),
		middleware.Logger,
	)) // route

	mux.Handle("POST /products",manager.With(
		http.HandlerFunc(handlers.CreateProduct),
		middleware.Logger,
	)) // route

	mux.Handle("GET /products/{productID}",manager.With(
		http.HandlerFunc(handlers.GetProductByID),
		middleware.Logger,
	)) // route

	fmt.Println("Server running on : 8080")

	err := http.ListenAndServe(":8080",global_router.GlobalRouter(mux))

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}