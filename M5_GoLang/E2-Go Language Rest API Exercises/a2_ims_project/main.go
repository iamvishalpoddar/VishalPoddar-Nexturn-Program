package main

import (
	db "InventoryManagementSystem/config"
	"InventoryManagementSystem/controller"
	"InventoryManagementSystem/middleware"
	"InventoryManagementSystem/repository"
	"InventoryManagementSystem/services"
	"fmt"
	"net/http"
)

func main() {
	db.InititalizeDB()

	productRepo := repository.NewProductRepository(db.GetDB())
	productService := services.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	productRoutes := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			productController.CreateProduct(w, r)
		case http.MethodGet:
			id := r.URL.Query().Get("id")
			fmt.Println("id received:", id)
			if id == "" {
				productController.GetAllProduct(w, r)
			} else {
				productController.GetProduct(w, r)
			}
		}
	})

	paramRoutes := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			productController.UpdateProduct(w, r)
		case http.MethodDelete:
			productController.DeleteProduct(w, r)
		}
	})

	protectedMux := http.NewServeMux()

	protectedMux.Handle("/product", productRoutes)
	protectedMux.Handle("/product/", paramRoutes)

	protectedRoutes := middleware.AuthorizationMiddleware(db.GetDB(), protectedMux)

	http.Handle("/product", protectedRoutes)
	http.Handle("/product/", protectedRoutes)

	loggedMux := middleware.LoggingMiddleware(http.DefaultServeMux)

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		fmt.Println("Server failed to start")
	}
}
