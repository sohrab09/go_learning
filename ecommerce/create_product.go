package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func createProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		fmt.Println(err)
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
