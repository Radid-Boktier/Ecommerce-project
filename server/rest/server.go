package rest

import (
	"ecommerce-server/config"
	"ecommerce-server/rest/handlers/product"
	"ecommerce-server/rest/handlers/review"
	"ecommerce-server/rest/handlers/user"
	middleware "ecommerce-server/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Server struct{
	cnf *config.Config
	productHandler *product.Handler
	userHandler *user.Handler
	reviewHandler *review.Handler
}

func NewServer(
	cnf *config.Config,
	productHandler *product.Handler,
	userHandler *user.Handler,
	reviewHandler *review.Handler,
) *Server {
	return &Server{
		cnf: cnf,
		productHandler: productHandler,
		userHandler: userHandler,
		reviewHandler: reviewHandler,
	}
}

func (server *Server) Start() {
	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux() // router
	wrappedMux := manager.WrapMux(mux)

	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)
	server.reviewHandler.RegisterRoutes(mux, manager)
	
	addr := ":" + strconv.Itoa(server.cnf.HttpPort)
	fmt.Println("Server running on port", addr)
	err := http.ListenAndServe(addr,wrappedMux)

	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}