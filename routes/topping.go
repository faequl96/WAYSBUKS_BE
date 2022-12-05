package routes

import (
	"waysbucks_api/handlers"
	"waysbucks_api/pkg/middleware"
	"waysbucks_api/pkg/mysql"
	"waysbucks_api/repositories"

	"github.com/gorilla/mux"
)

func ToppingRoutes(r *mux.Router) {
	toppingRepository := repositories.RepositoryTopping(mysql.DB)
	h := handlers.HandlerTopping(toppingRepository)

	r.HandleFunc("/toppings", h.HandlerGetToppings).Methods("GET")
	r.HandleFunc("/topping/{id}", h.HandlerGetToppingByID).Methods("GET")
	r.HandleFunc("/topping", middleware.Auth(middleware.UploadFile(h.HandlerCreateTopping))).Methods("POST")
	r.HandleFunc("/topping/{id}", middleware.Auth(middleware.UploadFile(h.HandlerUpdateTopping))).Methods("PATCH")
	r.HandleFunc("/topping/{id}", middleware.Auth(h.HandlerDeleteTopping)).Methods("DELETE")
}
