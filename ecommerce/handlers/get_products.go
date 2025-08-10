package handlers

import (
	"ecommerce/database"
	"encoding/json"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	response := map[string]any{
		"status":  http.StatusOK,
		"message": "Product list fetched successfully",
		"data":    database.ProductList,
	}
	json.NewEncoder(w).Encode(response)
}
