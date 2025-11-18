package rest

import (
	"ecommerce-server/rest/handlers"
	middleware "ecommerce-server/rest/middlewares"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager){
	

	mux.Handle("GET /products",manager.With(
		http.HandlerFunc(handlers.GetProducts),
	)) // route

	mux.Handle("POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreateProduct),
			middleware.AuthenticateJWT,
		),
	) // route

	mux.Handle("GET /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.GetProduct),
		),
	) // route

	mux.Handle("PUT /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.UpdateProduct),
			middleware.AuthenticateJWT,
		),
	) // route

	mux.Handle(
		"DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.DeleteProduct),
			middleware.AuthenticateJWT,
		),
	) // route

	mux.Handle(
		"POST /users",
		manager.With(
			http.HandlerFunc(handlers.CreateUser),
		),
	) // route

	mux.Handle(
		"POST /users/login",
		manager.With(
			http.HandlerFunc(handlers.Login),
		),
	) // route
}