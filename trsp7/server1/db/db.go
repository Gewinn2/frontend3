package db

import (
	"database/sql"
	"fmt"
	"log"
	"practice4/graph/model"

	"github.com/lib/pq"
)

func InitDataBase() (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		"postgres",
		"12345678",
		"db",
		"5432",
		"frontend_mirea",
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database connection: %v", err)
	}

	err = InitTables(db)
	if err != nil {
		log.Fatalf("Error initializing tables: %v", err)
	}

	return db, nil
}

func InitTables(db *sql.DB) error {
	createTableProducts := `
		CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			price DECIMAL(10, 2) NOT NULL,
			category TEXT[]
		);
	`

	_, err := db.Exec(createTableProducts)
	return err
}

func GetAllProducts(db *sql.DB) ([]*model.Product, error) {
	query := `
		SELECT id, name, price, category
		FROM products
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("GetAllProducts: %v", err)
	}
	defer rows.Close()

	var products []*model.Product
	for rows.Next() {
		var product model.Product
		var categories []string
		err := rows.Scan(&product.ID, &product.Name, &product.Price, pq.Array(&categories))
		if err != nil {
			return nil, fmt.Errorf("GetAllProducts: %v", err)
		}
		product.Category = categories
		products = append(products, &product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAllProducts: %v", err)
	}

	return products, nil
}
