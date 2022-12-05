package routes

import (
	"waysbucks_api/handlers"
	"waysbucks_api/pkg/middleware"
	"waysbucks_api/pkg/mysql"
	"waysbucks_api/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepoTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)
	r.HandleFunc("/create-transaction", middleware.Auth(h.HandlerCreateTransaction)).Methods("POST")
	r.HandleFunc("/transaction", middleware.Auth(h.HandlerGetTransactionByIdUser)).Methods("GET")
	r.HandleFunc("/transactions", middleware.Auth(h.HandlerGetTransactions)).Methods("GET")
	r.HandleFunc("/notification", h.Notification).Methods("POST")
}
