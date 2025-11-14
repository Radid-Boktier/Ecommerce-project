package cmd

import (
	"ecommerce-server/global_router"
	"ecommerce-server/handlers"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux() // router

	mux.Handle("GET /products",http.HandlerFunc(handlers.GetProducts))

	mux.Handle("POST /products",http.HandlerFunc(handlers.CreateProduct))

	mux.Handle("GET /products/{productID}",http.HandlerFunc(handlers.GetProductByID))

	fmt.Println("Server running on : 8080")

	err := http.ListenAndServe(":8080",global_router.GlobalRouter(mux))

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}