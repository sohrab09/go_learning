package main

import (
	"encoding/json"
	"net/http"
)

func getProducts(w http.ResponseWriter, r *http.Request) {

	response := map[string]any{
		"status":  http.StatusOK,
		"message": "Product list fetched successfully",
		"data":    productList,
	}
	json.NewEncoder(w).Encode(response)
}
