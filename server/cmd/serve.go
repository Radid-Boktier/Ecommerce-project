package cmd

import (
	"ecommerce-server/global_router"
	"ecommerce-server/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()

	manager.Use(middleware.Logger,middleware.Test1)

	mux := http.NewServeMux() // router

	initRoutes(mux,manager);

	fmt.Println("Server running on : 8080")

	err := http.ListenAndServe(":8080",global_router.GlobalRouter(mux))

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}