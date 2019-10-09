package api

import (
	"encoding/json"
	"net/http"

	"gopi.com/gopi/storage"
)

// GetProducts get products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.GetProducts())
}
