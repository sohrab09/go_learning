package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Product represents a product object
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imgUrl"`
}

// In-memory product store
var productList []Product

// Reusable function to set common headers
func setJSONHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
}

// Handlers
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the about page.")
}

func getAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	setJSONHeaders(w)
	/*
		// No need to check if the request method is GET or Not in new version of route.

		if r.Method != http.MethodGet {
			http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
			return
		}

	*/

	response := map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Product list fetched successfully",
		"data":    productList,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func createProductHandler(w http.ResponseWriter, r *http.Request) {
	setJSONHeaders(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

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

	response := map[string]interface{}{
		"status":  http.StatusCreated,
		"message": "Product created successfully",
		"data":    newProduct,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Initial seed data
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

	// Route setup

	// ?? Old Route setup

	mux := http.NewServeMux()
	// mux.HandleFunc("/hello", helloHandler)
	// mux.HandleFunc("/about", aboutHandler)
	// mux.HandleFunc("/products", getAllProductsHandler)
	// mux.HandleFunc("/create-product", createProductHandler)

	// !! New Route Setup

	mux.Handle("GET /hello", http.HandlerFunc(helloHandler))
	mux.Handle("GET /about", http.HandlerFunc(aboutHandler))
	mux.Handle("GET /products", http.HandlerFunc(getAllProductsHandler))
	mux.Handle("POST /create-product", http.HandlerFunc(createProductHandler))

	fmt.Println("🚀 Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("❌ Error starting server:", err)
	}
}

/*
	!! -> Name of some famous golang framework
		1. Gin
		2. Chi
		3. Gorilla Mux
*/
