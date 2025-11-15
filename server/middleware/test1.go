package middleware

import (
	"log"
	"net/http"
)

func Test1(next http.Handler) http.Handler {
	return  http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		log.Println("This is test1")

		next.ServeHTTP(w, r)
		
	})
}