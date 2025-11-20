package cmd

import (
	"ecommerce-server/config"
	"ecommerce-server/infra/db"
	"ecommerce-server/repo"
	"ecommerce-server/rest"
	"ecommerce-server/rest/handlers/product"
	"ecommerce-server/rest/handlers/review"
	"ecommerce-server/rest/handlers/user"
	middleware "ecommerce-server/rest/middlewares"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	productRpo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo(dbCon)
	reviewHandler := review.NewHandler()

	middlewares := middleware.NewMiddlewares(cnf)

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