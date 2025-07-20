package main

import (
	"database/sql"
	"example.com/cosmetics/config"
	"example.com/cosmetics/db"
	"example.com/cosmetics/handlers"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()

	if err := db.Init(cfg); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.DB.Close()

	r := chi.NewRouter()

	r.Route("/api/products", func(r chi.Router) {
		r.Post("/", handlers.CreateProduct(db.DB))
		r.Get("/", handlers.ListProducts(db.DB))
		r.Get("/{id}", handlers.GetProduct(db.DB))
		r.Put("/{id}", handlers.UpdateProduct(db.DB))
		r.Delete("/{id}", handlers.DeleteProduct(db.DB))
	})

	port := cfg.AppPort
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
