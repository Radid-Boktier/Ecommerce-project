package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I am a learner.")
}

type Product struct {
	ID int  `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	ImgUrl string `json:"imageUrl"`
}
 
var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request){
	handleCors(w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
	sendData(w,productList,200)
}

func createProduct(w http.ResponseWriter, r *http.Request){

	handleCors(w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	// r.Body => description,imageUrl,price,title => product er ekta instance => productList => append
	/*
		1.take body information (description,imageUrl,price,title) from r.body
		2.create an instance using Product struct with the body information
		3.append the instance into productList
	*/
	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid json", 400)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	sendData(w,newProduct,201)
}

func handleCors(w http.ResponseWriter){
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Methods","GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers","Content-Type, Boktier")
	w.Header().Set("Content-Type","application/json")
}

func handlePreflightReq(w http.ResponseWriter, r *http.Request){
	if r.Method == "OPTIONS" {
		w.WriteHeader(200) 
	}
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int){
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
func main() {
	mux := http.NewServeMux() // router

	mux.Handle("GET /hello", http.HandlerFunc((helloHandler))) // route
	// mux.HandleFunc("/hello",helloHandler)

	mux.Handle("GET /about",http.HandlerFunc(aboutHandler))

	mux.Handle("GET /products",http.HandlerFunc(getProducts))
	mux.Handle("OPTIONS /products",http.HandlerFunc(getProducts))

	mux.Handle("POST /create-products",http.HandlerFunc(createProduct))
	mux.Handle("OPTIONS /create-products",http.HandlerFunc(createProduct))
	
	fmt.Println("Server running on : 8080")

	err := http.ListenAndServe(":8080",mux)

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}

func init() { 
	prd1 := Product{
		ID: 1,
		Title: "Orange",
		Description: "Orange is red. I love orange",
		Price: 100,
		ImgUrl: "https://www.dole.com/sites/default/files/media/2025-01/oranges.png",
	}
	prd2 := Product{
		ID: 2,
		Title: "Apple",
		Description: "Apple is red. I love Apple",
		Price: 100,
		ImgUrl: "https://www.harrisfarm.com.au/cdn/shop/products/40715-done.jpg?v=1623908361&width=1946",
	}
	prd3 := Product{
		ID: 3,
		Title: "Banana",
		Description: "Banana is yellow. I love Banana",
		Price: 100,
		ImgUrl: "https://www.healthxchange.sg/adobe/dynamicmedia/deliver/dm-aid--bc117ef2-13c7-4c76-805d-3a4ad592a918/good-reasons-to-eat-a-banana-today.jpg?preferwebp=true",
	}
	// prd4 := Product{
	// 	ID: 1,
	// 	Title: "Grape",
	// 	Description: "Grape is green. I love grape",
	// 	Price: 200,
	// 	ImgUrl: "https://nationwideplants.com/cdn/shop/files/flame_seedless_grapes_vinyard.jpg?v=1731040443&width=1214",
	// }
	// prd5 := Product{
	// 	ID: 2,
	// 	Title: "Mango",
	// 	Description: "Mango is red. I love mango",
	// 	Price: 1000,
	// 	ImgUrl: "https://content.presspage.com/uploads/1460/69fbd9b4-d9fb-4591-8ede-0deb9c917a13/1920_stock-photo-fresh-mango-129537233.jpg?10000",
	// }
	// prd6 := Product{
	// 	ID: 3,
	// 	Title: "Pomegranate",
	// 	Description: "Pomegranate is red. I love Pomegranate",
	// 	Price: 300,
	// 	ImgUrl: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRlBvUi9p7zyfsUPbybYNRNeUN6rL5pT3-2Cg&s",
	// }

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	// productList = append(productList, prd4)
	// productList = append(productList, prd5)
	// productList = append(productList, prd6)
}