package routes

import (
	"waysbucks_api/handlers"
	"waysbucks_api/pkg/middleware"
	"waysbucks_api/pkg/mysql"
	"waysbucks_api/repositories"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {
	productRepository := repositories.RepositoryProduct(mysql.DB)
	h := handlers.HandlerProduct(productRepository)

	r.HandleFunc("/products", h.HandlerGetProducts).Methods("GET")
	r.HandleFunc("/product/{id}", h.HandlerGetProductByID).Methods("GET")
	r.HandleFunc("/product", middleware.Auth(middleware.UploadFile(h.HandlerCreateProduct))).Methods("POST")
	r.HandleFunc("/product/{id}", middleware.Auth(middleware.UploadFile(h.HandlerUpdateProduct))).Methods("PATCH")
	r.HandleFunc("/product/{id}", middleware.Auth(h.HandlerDeleteProduct)).Methods("DELETE")
}
