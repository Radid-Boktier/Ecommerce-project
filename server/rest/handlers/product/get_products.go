package product

import (
	"ecommerce-server/util"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request){
	productList, err := h.productRepo.List()
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server error")
		return
	}
	util.SendData(w, http.StatusOK, productList)
}