package database

import (
	"database/sql"
	"log"
)

// RunMigrations creates the base tables if they do not exist
func RunMigrations(DB *sql.DB) {
	tables := []string{

		// Staging table
		`CREATE TABLE IF NOT EXISTS staging (
            transaction_id VARCHAR(50),
            transaction_date DATE,
            user_id VARCHAR(50),
            country VARCHAR(100),
            region VARCHAR(100),
            product_id VARCHAR(50),
            product_name VARCHAR(255),
            category VARCHAR(100),
            price DECIMAL(10,2),
            quantity INT,
            total_price DECIMAL(12,2),
            stock_quantity INT,
            added_date DATE
        );`,

		// Users table
		`CREATE TABLE IF NOT EXISTS users (
			user_id VARCHAR(50) PRIMARY KEY,
			country VARCHAR(100),
			region VARCHAR(100)
		);`,

		// Products table
		`CREATE TABLE IF NOT EXISTS products (
			product_id VARCHAR(50) PRIMARY KEY,
			product_name VARCHAR(255),
			category VARCHAR(100),
			stock_quantity INT,
			added_date DATE
		);`,

		// Transactions table
		`CREATE TABLE IF NOT EXISTS transactions (
			transaction_id VARCHAR(50) PRIMARY KEY,
			transaction_date TIMESTAMP,
			user_id VARCHAR(50),
			product_id VARCHAR(50),
			price DECIMAL(10,2),
			quantity INT,
			total_price DECIMAL(12,2),
			CONSTRAINT fk_transaction_user FOREIGN KEY (user_id) REFERENCES users(user_id),
			CONSTRAINT fk_transaction_product FOREIGN KEY (product_id) REFERENCES products(product_id)
		);`,
	}

	for _, table := range tables {
		if _, err := DB.Exec(table); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
	}

	log.Println("Tables migration complete!")
}

// RunMaterializedViews creates materialized views safely after tables exist
func RunMaterializedViews(DB *sql.DB) {
	views := []string{
		// Top 20 products
		`CREATE MATERIALIZED VIEW IF NOT EXISTS top_products_mv AS
		 SELECT p.product_name,
		        SUM(t.quantity) AS items_sold,
		        p.stock_quantity
		 FROM transactions t
		 JOIN products p ON t.product_id = p.product_id
		 GROUP BY p.product_id, p.product_name, p.stock_quantity
		 ORDER BY items_sold DESC
		 LIMIT 20;`,

		// Country-level revenue
		`CREATE MATERIALIZED VIEW IF NOT EXISTS country_report_mv AS
		 SELECT u.country,
                ARRAY(SELECT DISTINCT jsonb_array_elements_text(jsonb_agg(p.product_name))) AS products,
		        SUM(t.total_price) AS total_revenue,
		        COUNT(DISTINCT t.transaction_id) AS transactions
		 FROM transactions t
		 JOIN users u ON t.user_id = u.user_id
		 JOIN products p ON t.product_id = p.product_id
		 GROUP BY u.country
		 ORDER BY total_revenue DESC;`,

		// Monthly sales
		`CREATE MATERIALIZED VIEW IF NOT EXISTS monthly_sales_mv AS
		 SELECT TO_CHAR(t.transaction_date, 'Mon') AS month,
		        DATE_PART('month', t.transaction_date) AS month_num,
		        SUM(t.quantity) AS units_sold
		 FROM transactions t
		 GROUP BY month, month_num
		 ORDER BY month_num;`,

		// Top 30 regions by revenue
		`CREATE MATERIALIZED VIEW IF NOT EXISTS regions_by_revenue_mv AS
		 SELECT u.region,
		        SUM(t.total_price) AS revenue_usd,
		        SUM(t.quantity) AS items_sold
		 FROM transactions t
		 JOIN users u ON t.user_id = u.user_id
		 GROUP BY u.region
		 ORDER BY revenue_usd DESC
		 LIMIT 30;`,
	}

	for _, view := range views {
		if _, err := DB.Exec(view); err != nil {
			log.Fatalf("Materialized view creation failed: %v", err)
		}
	}

	log.Println("Materialized views created successfully!")
}

// RefreshMaterializedViews refreshes all materialized views (call periodically after new data)
func RefreshMaterializedViews(DB *sql.DB) {
	views := []string{
		"top_products_mv",
		"country_report_mv",
		"monthly_sales_mv",
		"regions_by_revenue_mv",
	}

	for _, view := range views {
		if _, err := DB.Exec("REFRESH MATERIALIZED VIEW " + view + ";"); err != nil {
			log.Fatalf("Refreshing materialized view %s failed: %v", view, err)
		}
	}

	log.Println("Materialized views refreshed successfully!")
}
