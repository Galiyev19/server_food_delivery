package service

import (
	"food_delivery/internal/storage"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductService struct {
	store storage.Storage
}

func NewProductService(s storage.Storage) *ProductService {
	return &ProductService{store: s}
}

func (s *ProductService) RegisterRoutes(r *httprouter.Router) {
	r.HandlerFunc(http.MethodGet, "/v1/products", s.handleCreateProduct)
	r.HandlerFunc(http.MethodGet, "/v1/products/:id", s.handleGetProduct)
}

func (s *ProductService) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
}

func (s *ProductService) handleGetProduct(w http.ResponseWriter, r *http.Request) {
}
