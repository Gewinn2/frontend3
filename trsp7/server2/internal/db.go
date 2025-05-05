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

	createMocks(db)

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

func createMocks(db *sql.DB) {
	query := `SELECT COUNT(*) FROM products;`

	count := 0
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count == 0 {
		query = `INSERT INTO products(name, price, category) VALUES ($1, $2, $3)`
		products := []Product{
			{
				Id:       1,
				Name:     "Ноутбук Lenovo IdeaPad",
				Price:    54999.99,
				Category: []string{"Электроника", "Ноутбуки"},
			},
			{
				Id:       2,
				Name:     "Смартфон iPhone 15",
				Price:    89999.99,
				Category: []string{"Электроника", "Смартфоны"},
			},
			{
				Id:       3,
				Name:     "Наушники Sony WH-1000XM5",
				Price:    29999.99,
				Category: []string{"Электроника", "Аксессуары"},
			},
			{
				Id:       4,
				Name:     "Кофемашина De'Longhi",
				Price:    45999.99,
				Category: []string{"Бытовая техника", "Кухня"},
			},
			{
				Id:       5,
				Name:     "Фитнес-браслет Xiaomi Mi Band 7",
				Price:    3999.99,
				Category: []string{"Гаджеты", "Спорт"},
			},
			{
				Id:       6,
				Name:     "Книга 'Чистый код' Роберт Мартин",
				Price:    2499.99,
				Category: []string{"Книги", "Программирование"},
			},
			{
				Id:       7,
				Name:     "Игровая консоль PlayStation 5",
				Price:    64999.99,
				Category: []string{"Игры", "Консоли"},
			},
			{
				Id:       8,
				Name:     "Беспроводная мышь Logitech MX Master 3",
				Price:    8999.99,
				Category: []string{"Компьютеры", "Аксессуары"},
			},
			{
				Id:       9,
				Name:     "Умная колонка Яндекс Станция 2",
				Price:    12999.99,
				Category: []string{"Умный дом", "Аудио"},
			},
			{
				Id:       10,
				Name:     "Электросамокат Xiaomi Pro 2",
				Price:    34999.99,
				Category: []string{"Транспорт", "Гаджеты"},
			},
		}
		for _, product := range products {
			_, err = CreateProduct(db, product)
			if err != nil {
				log.Println(err)
			}
		}
	}
	log.Println("моки успешно созданы")
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
