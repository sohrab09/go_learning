package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imgUrl"`
}

// In-memory product store
var productList []Product

func getAllProductsHandler(w http.ResponseWriter, r *http.Request) {

	response := map[string]any{
		"status":  http.StatusOK,
		"message": "Product list fetched successfully",
		"data":    productList,
	}
	json.NewEncoder(w).Encode(response)
}

func createProductHandler(w http.ResponseWriter, r *http.Request) {

	var newProduct Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Validation
	if newProduct.Name == "" || newProduct.Description == "" || newProduct.Price <= 0 || newProduct.ImgURL == "" {
		http.Error(w, "All fields (name, description, price, imgUrl) are required and must be valid", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	response := map[string]any{
		"status":  201,
		"message": "Product created successfully",
		"data":    newProduct,
	}
	json.NewEncoder(w).Encode(response)
}

func main() {

	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(getAllProductsHandler))
	mux.Handle("POST /create-product", http.HandlerFunc(createProductHandler))

	globalRouter := globalRouter(mux)

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

func globalRouter(mux *http.ServeMux) http.Handler {
	handleAllReq := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		w.WriteHeader(200)

		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
		mux.ServeHTTP(w, r)
	}
	return http.HandlerFunc(handleAllReq)
}
