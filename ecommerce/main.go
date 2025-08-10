package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	globalRouter := globalRouter(mux)

	mux.Handle("GET /products", http.HandlerFunc(getProducts))
	mux.Handle("POST /create-product", http.HandlerFunc(createProduct))

	fmt.Println("🚀 Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", globalRouter); err != nil {
		fmt.Println("❌ Error starting server:", err)
	}
}

func init() {
	productList = append(productList, Product{
		ID:          1,
		Name:        "Product 1",
		Description: "Description of product 1",
		Price:       10.99,
		ImgURL:      "https://example.com/product1.jpg",
	}, Product{
		ID:          2,
		Name:        "Product 2",
		Description: "Description of product 2",
		Price:       19.99,
		ImgURL:      "https://example.com/product2.jpg",
	})
}

/*
	১- যদি সব গুলো ফাইল একই ফোল্ডারে থাকে তাহলে এপ রান করার কমান্ড হলো -> go run .
	২- এবং প্যাকেজের নাম হবে main, কারন সব গুলো ফাইল একই প্যাকেজে থাকে।
*/
