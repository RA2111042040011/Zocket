package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	connStr := "postgres://user:password@localhost/productsdb?sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		product_name TEXT NOT NULL,
		product_description TEXT,
		product_images TEXT[],
		compressed_product_images TEXT[],
		product_price DECIMAL
	);`
	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatalf("Failed to create products table: %v", err)
	}
}
