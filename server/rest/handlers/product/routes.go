package product

import (
	middleware "ecommerce-server/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager){
	
	mux.Handle("GET /products",manager.With(
		http.HandlerFunc(h.GetProducts),
	)) // route

	mux.Handle("POST /products",
		manager.With(
			http.HandlerFunc(h.CreateProduct),
			middleware.AuthenticateJWT,
		),
	) // route

	mux.Handle("GET /products/{id}",
		manager.With(
			http.HandlerFunc(h.GetProduct),
		),
	) // route

	mux.Handle("PUT /products/{id}",
		manager.With(
			http.HandlerFunc(h.UpdateProduct),
			middleware.AuthenticateJWT,
		),
	) // route

	mux.Handle(
		"DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(h.DeleteProduct),
			middleware.AuthenticateJWT,
		),
	) // route

}