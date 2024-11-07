package controller

import (
	"encoding/json"
	"net/http"
	"product_manage/model"
	"product_manage/service"

	"github.com/gorilla/mux"
)

type ProductController struct {
	service      *service.ProductService
	wsController *WebSocketController
}

func NewProductController(repo *service.ProductService, wsController *WebSocketController) *ProductController {
	return &ProductController{
		service:      repo,
		wsController: wsController,
	}
}

func (pc *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if err := pc.service.CreateProduct(&product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// pc.wsController.NotifyProductChange(product.ID, "created")
	pc.wsController.SendProductChange(&product, "created")
	w.WriteHeader(http.StatusCreated)
}
func (pc *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if err := pc.service.UpdateProduct(&product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pc.wsController.NotifyProductChange(product.ID, "updated")
	pc.wsController.SendProductChange(&product, "updated")
	w.WriteHeader(http.StatusCreated)
}

func (pc *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	productID := product.ID
	if err := pc.service.DeleteProduct(productID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pc.wsController.NotifyProductChange(product.ID, "deleted")
	pc.wsController.SendProductChange(product.ID, "deleted")
	w.WriteHeader(http.StatusCreated)
}

func (pc *ProductController) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	var products []model.Product
	products, err := pc.service.GetAllProducts()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func (pc *ProductController) GetProductById(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	vars := mux.Vars(r)
	productID := vars["id"]
	product, err := pc.service.GetProductById(productID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
