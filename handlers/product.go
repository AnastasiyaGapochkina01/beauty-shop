package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"example.com/cosmetics/models"
	"github.com/go-chi/chi/v5"
)

func CreateProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.CreateProductRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		query := `INSERT INTO products (name, description, price, brand, category, stock_quantity) 
                  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
		var id int
		err := db.QueryRow(query, 
			req.Name, req.Description, req.Price, req.Brand, req.Category, req.StockQuantity).Scan(&id)
		if err != nil {
			http.Error(w, "Failed to create product: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]int{"id": id})
	}
}

func GetProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		var p models.Product
		query := `SELECT id, name, description, price, brand, category, stock_quantity, created_at, updated_at 
                  FROM products WHERE id = $1`
		err = db.QueryRow(query, id).Scan(
			&p.ID, &p.Name, &p.Description, &p.Price, &p.Brand, &p.Category, &p.StockQuantity, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				http.Error(w, "Product not found", http.StatusNotFound)
				return
			}
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)
	}
}

func UpdateProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		var req models.UpdateProductRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		query := `UPDATE products SET 
                  name = COALESCE($1, name),
                  description = COALESCE($2, description),
                  price = COALESCE($3, price),
                  brand = COALESCE($4, brand),
                  category = COALESCE($5, category),
                  stock_quantity = COALESCE($6, stock_quantity),
                  updated_at = CURRENT_TIMESTAMP
                  WHERE id = $7`
		res, err := db.Exec(query, 
			req.Name, req.Description, req.Price, req.Brand, req.Category, req.StockQuantity, id)
		if err != nil {
			http.Error(w, "Update failed: "+err.Error(), http.StatusInternalServerError)
			return
		}

		rows, _ := res.RowsAffected()
		if rows == 0 {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func DeleteProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		query := "DELETE FROM products WHERE id = $1"
		res, err := db.Exec(query, id)
		if err != nil {
			http.Error(w, "Delete failed: "+err.Error(), http.StatusInternalServerError)
			return
		}

		rows, _ := res.RowsAffected()
		if rows == 0 {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func ListProducts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name, description, price, brand, category, stock_quantity FROM products")
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var products []models.Product
		for rows.Next() {
			var p models.Product
			if err := rows.Scan(
				&p.ID, &p.Name, &p.Description, &p.Price, &p.Brand, &p.Category, &p.StockQuantity); err != nil {
				http.Error(w, "Scan error", http.StatusInternalServerError)
				return
			}
			products = append(products, p)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}
}
