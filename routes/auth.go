package routes

import (
	"waysbucks_api/handlers"
	"waysbucks_api/pkg/middleware"
	"waysbucks_api/pkg/mysql"
	"waysbucks_api/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerAuth(userRepository)

	r.HandleFunc("/register", h.HandlerRegister).Methods("POST")
	r.HandleFunc("/login", h.HandlerLogin).Methods("POST")
	r.HandleFunc("/check-auth", middleware.Auth(h.HandlerCheckAuth)).Methods("GET")
}
