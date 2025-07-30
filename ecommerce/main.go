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

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	if r.Method == "OPTIONS" {
		w.WriteHeader(201)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Plz give me POST request", 400)
		return
	}

	var newProduct Products
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		http.Error(w, "Please give me valid JSON", 400)
		return
	}

	// validation
	if newProduct.Name == "" {
		http.Error(w, "Please provide product name", 400)
		return
	} else if newProduct.Description == "" {
		http.Error(w, "Please provide product description", 400)
		return
	} else if newProduct.Price == 0 {
		http.Error(w, "Please provide product price", 400)
		return
	} else if newProduct.ImgUrl == "" {
		http.Error(w, "Please provide product image URL", 400)
		return
	} else {
		fmt.Println("Product created successfully")
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	encoder := json.NewEncoder(w)
	// encoder.Encode(newProduct)

	response := map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Product created successfully",
		"data":    newProduct,
	}

	encoder.Encode(response)

}

func main() {
	mux := http.NewServeMux() // in here "mux" is a router

	mux.HandleFunc("/hello", helloHandler) // in here "hello" is a pattern or route and "helloHandler" is a handler function that processes requests to that route or process something

	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/products", getAllProducts)
	mux.HandleFunc("/create-product", createProduct)

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
