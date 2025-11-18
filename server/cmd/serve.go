package cmd

import (
	"ecommerce-server/config"
	"ecommerce-server/rest"
	"ecommerce-server/rest/handlers/product"
	"ecommerce-server/rest/handlers/review"
	"ecommerce-server/rest/handlers/user"
	middleware "ecommerce-server/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	middlewares := middleware.NewMiddlewares(cnf)
	
	productHandler := product.NewHandler(middlewares)
	userHandler := user.NewHandler()
	reviewHandler := review.NewHandler()

	server := rest.NewServer(
		cnf,
		productHandler, 
		userHandler,
		reviewHandler,
	)
	server.Start()
}