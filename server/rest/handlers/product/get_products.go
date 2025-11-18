package product

import (
	"ecommerce-server/database"
	"ecommerce-server/util"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request){
	util.SendData(w,database.List(),200)
}