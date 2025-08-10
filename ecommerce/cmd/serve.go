package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()
	globalRouter := global_router.GlobalRouter(mux)

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductsByID))

	fmt.Println("🚀 Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", globalRouter); err != nil {
		fmt.Println("❌ Error starting server:", err)
	}
}
