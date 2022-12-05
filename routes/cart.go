package routes

import (
	"waysbucks_api/handlers"
	"waysbucks_api/pkg/middleware"
	"waysbucks_api/pkg/mysql"
	"waysbucks_api/repositories"

	"github.com/gorilla/mux"
)

func CartRoutes(r *mux.Router) {
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(cartRepository)
	r.HandleFunc("/add-cart", middleware.Auth(h.HandlerAddCart)).Methods("POST")
	r.HandleFunc("/cart/{id}", middleware.Auth(h.HandlerDeleteCart)).Methods("DELETE")
	r.HandleFunc("/cart", middleware.Auth(h.HandlerGetCartByIdUserIsPayedFalse)).Methods("GET")
	r.HandleFunc("/update-cart", middleware.Auth(h.HandlerUpdateCart)).Methods("PATCH")
}
