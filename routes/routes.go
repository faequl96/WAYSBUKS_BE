package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	AuthRoutes(r)
	UserRoutes(r)
	ProductRoutes(r)
	ToppingRoutes(r)
	CartRoutes(r)
	TransactionRoutes(r)
}
