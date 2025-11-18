package product

import (
	"ecommerce-server/database"
	"ecommerce-server/util"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct (w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "please give me a valid product id", 400)
		return
	}

	database.Delete(pId)

	util.SendData(w, "Successfully deleted product", 201)
}