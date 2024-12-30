package controller

import (
	"InventoryManagementSystem/model"
	"InventoryManagementSystem/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type ProductController struct {
	ProductService *services.ProductService
}

func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{service}
}

func (controller *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var product *model.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request body"+err.Error(), http.StatusBadRequest)
		return
	}

	createdProduct, err := controller.ProductService.CreateProduct(product)

	if err != nil {
		http.Error(w, "Error creating the product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(createdProduct); err != nil {
		http.Error(w, "Error encoding the product: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (controller *ProductController) GetProduct(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	productId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	product, err := controller.ProductService.GetProduct(productId)

	if err != nil {
		http.Error(w, "Error getting the product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(product); err != nil {
		http.Error(w, "Error encoding the product: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (controller *ProductController) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	products, err := controller.ProductService.GetAllProducts()

	if err != nil {
		http.Error(w, "Error getting the products: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "Error encoding the products: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (controller *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")

	productId, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	var product *model.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request body"+err.Error(), http.StatusBadRequest)
		return
	}
	product.ID = productId
	updatedProduct, err := controller.ProductService.UpdateProduct(product)

	if err != nil {
		http.Error(w, "Error updating the product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(updatedProduct); err != nil {
		http.Error(w, "Error encoding the product: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (controller *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	productId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	err = controller.ProductService.DeleteProduct(productId)

	if err != nil {
		http.Error(w, "Error deleting the product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode("Product deleted successfully"); err != nil {
		http.Error(w, "Error encoding the product: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
