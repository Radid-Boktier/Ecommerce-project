package middleware

import (
	"log"
	"net/http"
)

func Test2(next http.Handler) http.Handler {
	return  http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		log.Println("This is test2 ")
		
		next.ServeHTTP(w, r)
		
	})
}