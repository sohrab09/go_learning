package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the about page.")
}

type Products struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

var productList []Products

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Plz give me GET request", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Product list fetched successfully",
		"data":    productList,
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	mux := http.NewServeMux() // in here "mux" is a router

	mux.HandleFunc("/hello", helloHandler) // in here "hello" is a pattern or route and "helloHandler" is a handler function that processes requests to that route or process something

	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/products", getAllProducts)

	fmt.Println("Server is running on port 8080")

	http.ListenAndServe(":8080", mux)

	err := http.ListenAndServe(":8080", mux) // failed to start server

	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}

func init() {
	prd1 := Products{
		ID:          1,
		Name:        "Product 1",
		Description: "Description of product 1",
		Price:       10.99,
		ImgUrl:      "https://example.com/product1.jpg",
	}

	prd2 := Products{
		ID:          2,
		Name:        "Product 2",
		Description: "Description of product 2",
		Price:       19.99,
		ImgUrl:      "https://example.com/product2.jpg",
	}

	productList = append(productList, prd1, prd2)
}
