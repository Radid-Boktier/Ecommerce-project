package product

import (
	"ecommerce-server/util"
	"net/http"
	"strconv"
)


func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request){
	// get query parameters
	// get page from query params
	// get limit from query params
	reqQuery := r.URL.Query()

	pageAsString := reqQuery.Get("page")
	limitAsString := reqQuery.Get("limit")

	page, _ := strconv.ParseInt(pageAsString, 10, 32)
	limit, _ := strconv.ParseInt(limitAsString, 10, 32)
	
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	productList, err := h.svc.List(page, limit)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server error")
		return
	}

	cnt, err := h.svc.Count()
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server error")
		return
	}

	util.SendPage(w,productList, page, limit,cnt)
}