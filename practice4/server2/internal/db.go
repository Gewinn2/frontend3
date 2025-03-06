package internal

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
)

func InitDataBase() (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		"postgres",
		"Gew1234",
		"localhost",
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

func CreateProduct(db *sql.DB, product Product) (int, error) {
	query := `
		INSERT INTO products (name, price, category)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	var id int
	err := db.QueryRow(query, product.Name, product.Price, pq.Array(product.Category)).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("CreateProduct: %v", err)
	}

	return id, nil
}

func UpdateProduct(db *sql.DB, product Product) error {
	query := `
		UPDATE products
		SET name = $1, price = $2, category = $3
		WHERE id = $4
	`

	_, err := db.Exec(query, product.Name, product.Price, pq.Array(product.Category), product.Id)
	if err != nil {
		return fmt.Errorf("UpdateProduct: %v", err)
	}

	return nil
}

func DeleteProduct(db *sql.DB, id int) error {
	query := `
		DELETE FROM products
		WHERE id = $1
	`

	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("DeleteProduct: %v", err)
	}

	return nil
}
