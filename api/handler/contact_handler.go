package handler

import (
	"QCaller/api"
	"QCaller/types"
	"encoding/json"
	"net/http"
)

// CreateContact : handler for create contact api
func CreateContact(w http.ResponseWriter, r *http.Request) {
	ctx := types.ContextFrom(r)
	response := api.CreateContact(r, ctx)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetContact : handler for get contact api
func GetContact(w http.ResponseWriter, r *http.Request) {
	ctx := types.ContextFrom(r)
	response := api.GetContact(r, ctx)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// SearchContact : handler for search contact api
func SearchContact(w http.ResponseWriter, r *http.Request) {
	ctx := types.ContextFrom(r)
	response := api.SearchContact(r, ctx)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UpdateContact : handler for update contact api
func UpdateContact(w http.ResponseWriter, r *http.Request) {
	ctx := types.ContextFrom(r)
	response := api.UpdateContact(r, ctx)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteContact : handler for delete contact api
func DeleteContact(w http.ResponseWriter, r *http.Request) {
	ctx := types.ContextFrom(r)
	response := api.DeleteContact(r, ctx)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// BulkCreateContact : handler for bulk create contact api
func BulkCreateContact(w http.ResponseWriter, r *http.Request) {
	ctx := types.ContextFrom(r)
	response := api.BulkCreateContact(r, ctx)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// BulkUpdateContact : handler for bulk update contact api
func BulkUpdateContact(w http.ResponseWriter, r *http.Request) {
	ctx := types.ContextFrom(r)
	response := api.BulkUpdateContact(r, ctx)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
