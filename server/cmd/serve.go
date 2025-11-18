package cmd

import (
	"ecommerce-server/config"
	"ecommerce-server/rest"
	"ecommerce-server/rest/handlers/product"
	"ecommerce-server/rest/handlers/review"
	"ecommerce-server/rest/handlers/user"
)

func Serve() {
	cnf := config.GetConfig()
	productHandler := product.NewHandler()
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