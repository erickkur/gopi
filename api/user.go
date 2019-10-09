package api

import (
	"encoding/json"
	"net/http"

	"gopi.com/gopi/storage"
)

// GetUsers get users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.GetUsers())
}
