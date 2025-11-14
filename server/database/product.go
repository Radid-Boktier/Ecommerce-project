package database

var ProductList []Product

type Product struct {
	ID int  `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	ImgUrl string `json:"imageUrl"`
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
	prd4 := Product{
		ID: 4,
		Title: "Grape",
		Description: "Grape is green. I love grape",
		Price: 200,
		ImgUrl: "https://nationwideplants.com/cdn/shop/files/flame_seedless_grapes_vinyard.jpg?v=1731040443&width=1214",
	}
	prd5 := Product{
		ID: 5,
		Title: "Mango",
		Description: "Mango is red. I love mango",
		Price: 1000,
		ImgUrl: "https://content.presspage.com/uploads/1460/69fbd9b4-d9fb-4591-8ede-0deb9c917a13/1920_stock-photo-fresh-mango-129537233.jpg?10000",
	}
	prd6 := Product{
		ID: 6,
		Title: "Pomegranate",
		Description: "Pomegranate is red. I love Pomegranate",
		Price: 300,
		ImgUrl: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRlBvUi9p7zyfsUPbybYNRNeUN6rL5pT3-2Cg&s",
	}

	ProductList = append(ProductList, prd1)
	ProductList = append(ProductList, prd2)
	ProductList = append(ProductList, prd3)
	ProductList = append(ProductList, prd4)
	ProductList = append(ProductList, prd5)
	ProductList = append(ProductList, prd6)
}