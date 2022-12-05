package routes

import (
	"waysbucks_api/handlers"
	"waysbucks_api/pkg/middleware"
	"waysbucks_api/pkg/mysql"
	"waysbucks_api/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/users", middleware.Auth(h.HandlerGetUsers)).Methods("GET")
	r.HandleFunc("/user/{id}", middleware.Auth(h.HandlerGetUserByID)).Methods("GET")
	r.HandleFunc("/user/{id}", middleware.Auth(middleware.UploadFile(h.HandlerUpdateUser))).Methods("PATCH")
	r.HandleFunc("/user/{id}", middleware.Auth(h.HandlerDeleteUser)).Methods("DELETE")
}
