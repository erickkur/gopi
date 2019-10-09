package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gopi.com/gopi/api"
	"gopi.com/gopi/migration"
	"gopi.com/gopi/storage"
)

func main() {
	// Prepare database
	storage.PrepareDB()
	defer storage.GopiDB.Close()

	// Migration
	args := os.Args[1:]
	if len(args) > 0 {
		switch args[0] {
		case "migrate":
			m := new(migration.Migration)
			m.RunMigrations()
			return
		}
	}

	// Run server
	r := mux.NewRouter()
	r.HandleFunc("/users", api.GetUsers).Methods("GET")
	r.HandleFunc("/products", api.GetProducts).Methods("GET")
	log.Print("listen port 3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
