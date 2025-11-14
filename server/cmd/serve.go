package cmd

import (
	"ecommerce-server/global_router"
	"ecommerce-server/handlers"
	"ecommerce-server/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux() // router

	mux.Handle("GET /route", middleware.Logger(http.HandlerFunc(handlers.Test)))

	mux.Handle("GET /products",middleware.Logger(http.HandlerFunc(handlers.GetProducts))) // route

	mux.Handle("POST /products",middleware.Logger(http.HandlerFunc(handlers.CreateProduct))) // route

	mux.Handle("GET /products/{productID}",middleware.Logger(http.HandlerFunc(handlers.GetProductByID))) // route

	fmt.Println("Server running on : 8080")

	err := http.ListenAndServe(":8080",global_router.GlobalRouter(mux))

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}