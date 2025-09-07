package database

import (
	"database/sql"
	"log"
)

func RunMigrations(DB *sql.DB) {
	tables := []string{
		// Users table
		`CREATE TABLE IF NOT EXISTS users (
			user_id INT PRIMARY KEY,
			country VARCHAR(100),
			region VARCHAR(100)
		);`,

		// Products table
		`CREATE TABLE IF NOT EXISTS products (
			product_id INT PRIMARY KEY,
			product_name VARCHAR(255),
			category VARCHAR(100),
			price DECIMAL(10,2),
			stock_quantity INT,
			added_date TIMESTAMP
		);`,

		// Transactions table
		`CREATE TABLE IF NOT EXISTS transactions (
			transaction_id INT PRIMARY KEY,
			transaction_date TIMESTAMP,
			user_id INT,
			CONSTRAINT fk_transaction_user FOREIGN KEY (user_id) REFERENCES users(user_id)
		);`,

		// Transaction details table
		`CREATE TABLE IF NOT EXISTS transaction_details (
			transaction_id INT,
			product_id INT,
			quantity INT,
			total_price DECIMAL(10,2),
			PRIMARY KEY (transaction_id, product_id),
			CONSTRAINT fk_detail_transaction FOREIGN KEY (transaction_id) REFERENCES transactions(transaction_id),
			CONSTRAINT fk_detail_product FOREIGN KEY (product_id) REFERENCES products(product_id)
		);`,
	}

	for _, table := range tables {
		if _, err := DB.Exec(table); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
	}

	log.Println("Migrations complete!")
}
