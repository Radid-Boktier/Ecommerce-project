package cmd

import (
	"ecommerce-server/config"
	"ecommerce-server/infra/db"
	"ecommerce-server/product"
	"ecommerce-server/repo"
	"ecommerce-server/rest"
	prodHandler "ecommerce-server/rest/handlers/product"
	"ecommerce-server/rest/handlers/review"
	usrHandler "ecommerce-server/rest/handlers/user"
	middleware "ecommerce-server/rest/middlewares"
	"ecommerce-server/user"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	//repos
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)
	reviewHandler := review.NewHandler()

	//domains
	usrSvc := user.NewService(userRepo)
	prdctSvc := product.NerService(productRepo)

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := prodHandler.NewHandler(middlewares,prdctSvc)
	userHandler := usrHandler.NewHandler(cnf, usrSvc)

	server := rest.NewServer(
		cnf,
		productHandler, 
		userHandler,
		reviewHandler,
	)
	server.Start()
}