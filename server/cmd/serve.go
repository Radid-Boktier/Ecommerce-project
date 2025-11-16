package cmd

import (
	"ecommerce-server/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux() // router
	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux,manager);

	fmt.Println("Server running on : 8080")

	err := http.ListenAndServe(":8080",wrappedMux)

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}