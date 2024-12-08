package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"product-management/db"
	"product-management/handlers"
)

func main() {
	db.InitDB()
	r := mux.NewRouter()

	r.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", handlers.GetProductByID).Methods("GET")

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}
