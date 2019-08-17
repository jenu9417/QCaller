package routes

import (
	"QCaller/api/handler"

	"github.com/gorilla/mux"
)

// NewRouter : returns instance of Router which routes API requests to the corresponding handler func
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/contact", handler.GetContact).Methods("GET")
	router.HandleFunc("/contact", handler.CreateContact).Methods("POST")
	router.HandleFunc("/contact", handler.UpdateContact).Methods("PUT")
	router.HandleFunc("/contact", handler.DeleteContact).Methods("DELETE")
	router.HandleFunc("/contact/bulk", handler.BulkCreateContact).Methods("POST")
	router.HandleFunc("/contact/bulk", handler.BulkUpdateContact).Methods("PUT")
	router.HandleFunc("/contact/search", handler.SearchContact).Methods("GET")

	return router
}
