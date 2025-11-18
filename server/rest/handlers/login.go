package handlers

import (
	"ecommerce-server/database"
	"ecommerce-server/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginUser struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request){
	var loginUser database.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginUser)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	usr := database.Find(loginUser.Email, loginUser.Password)

	if(usr == nil) {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
	}
	util.SendData(w,usr,http.StatusCreated)
}