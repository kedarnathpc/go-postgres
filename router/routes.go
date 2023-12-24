package router

import (
	"github.com/gorilla/mux"
	"github.com/kedarnathpc/go-postgres/middleware"
)

// Router returns a new instance of the Gorilla Mux router with configured routes for stock-related operations.
func Router() *mux.Router {
	// Create a new instance of the Gorilla Mux router
	router := mux.NewRouter()

	// Define routes for various stock operations and associate them with corresponding middleware functions
	router.HandleFunc("/api/stock/{id}", middleware.GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stock", middleware.GetAllStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newstock", middleware.CreateStock).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/stock/{id}", middleware.UpdateStock).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deletestock/{id}", middleware.DeleteStock).Methods("DELETE", "OPTIONS")

	// Return the configured router
	return router
}
