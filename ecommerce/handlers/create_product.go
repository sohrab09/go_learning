package handlers

import (
	"ecommerce/database" // this is the way to import the specific package from different folder
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product
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

	newProduct.ID = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, newProduct)

	response := map[string]any{
		"status":  201,
		"message": "Product created successfully",
		"data":    newProduct,
	}
	json.NewEncoder(w).Encode(response)
}
