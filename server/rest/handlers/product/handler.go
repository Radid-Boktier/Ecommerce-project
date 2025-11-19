package product

import (
	"ecommerce-server/repo"
	middleware "ecommerce-server/rest/middlewares"
)

type Handler struct {
	middlewares *middleware.Middlewares
	productRepo repo.ProductRepo
}

func NewHandler(
	middlewares *middleware.Middlewares,
	productRepo repo.ProductRepo,
) *Handler {
	return  &Handler{
		middlewares: middlewares,
		productRepo:productRepo,
	}
}