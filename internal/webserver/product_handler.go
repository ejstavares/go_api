package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/ejstavares/goapi/internal/entity"
	"github.com/ejstavares/goapi/internal/services"
	"github.com/go-chi/chi/v5"
)

type WebProductHandler struct {
	ProductService *services.ProductService
}

func NewWebProductHandler(productService services.ProductService) *WebProductHandler  {
	return &WebProductHandler{ProductService: &productService}
}


func (wch *WebProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := wch.ProductService.GetProducts()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

func (wch *WebProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	 err := json.NewDecoder(r.Body).Decode(&product)

	 if err != nil {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	 }

	result, err := wch.ProductService.CreateProduct(product.Name, product.Description,product.Price, product.CategoryID, product.ImageURL)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (wch *WebProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r,"id")

	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	product, err := wch.ProductService.GetProduct(id)

	if err!= nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	json.NewEncoder(w).Encode(product)
}
func (wch *WebProductHandler) GetProductByCategoryID(w http.ResponseWriter, r *http.Request) {
	categoryID := chi.URLParam(r,"categoryID")

	if categoryID == "" {
		http.Error(w, "categoryID is required", http.StatusBadRequest)
		return
	}

	product, err := wch.ProductService.GetProductByCategoryID(categoryID)

	if err!= nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	json.NewEncoder(w).Encode(product)
}