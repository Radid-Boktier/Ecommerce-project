package cmd

import (
	"ecommerce-server/config"
	"ecommerce-server/repo"
	"ecommerce-server/rest"
	"ecommerce-server/rest/handlers/product"
	"ecommerce-server/rest/handlers/review"
	"ecommerce-server/rest/handlers/user"
	middleware "ecommerce-server/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	middlewares := middleware.NewMiddlewares(cnf)
	
	productRpo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo()
	reviewHandler := review.NewHandler()

	productHandler := product.NewHandler(middlewares,productRpo)

	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(
		cnf,
		productHandler, 
		userHandler,
		reviewHandler,
	)
	server.Start()
}