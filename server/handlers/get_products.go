package handlers

import (
	"ecommerce-server/database"
	"ecommerce-server/util"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request){
	util.SendData(w,database.ProductList,200)
}