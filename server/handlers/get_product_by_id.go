package handlers

import (
	"ecommerce-server/database"
	"ecommerce-server/util"
	"net/http"
	"strconv"
)



func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue(("productID"))

	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "please give me a valid product id", 400)
		return
	}

	for _, product := range database.ProductList {
		if product.ID == pId {
			util.SendData(w, product, 200)
			return
		}
	}

	util.SendData(w, "Product not found", 404)
}