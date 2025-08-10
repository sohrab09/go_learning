package handlers

import (
	"ecommerce/database"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetProductsByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	// convert string to int
	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Invalid product ID", 400)
		return
	}

	// find the product by ID
	for _, product := range database.ProductList {
		if product.ID == pId {
			response := map[string]any{
				"status":  http.StatusOK,
				"message": "Product fetched successfully",
				"data":    database.ProductList[pId-1],
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}
