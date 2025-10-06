package database

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
)

// LoadCSVAndSync streams CSV into staging table and normalizes
func LoadCSVAndSync(conn *pgx.Conn, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Skip header
	if _, err := reader.Read(); err != nil {
		return err
	}

	batch := make([][]interface{}, 0, 100_000) // batch size
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return err
		}

		transactionDate, err := time.Parse("2006-01-02", record[12])
		if err != nil {
			return fmt.Errorf("invalid date in CSV: %q, err: %w", record[12], err)
		}
		price, _ := strconv.ParseFloat(record[8], 64)
		quantity, _ := strconv.Atoi(record[9])
		totalPrice, _ := strconv.ParseFloat(record[10], 64)
		stockQuantity, _ := strconv.Atoi(record[11])
		addedDate, err := time.Parse("2006-01-02", record[12])
		if err != nil {
			return fmt.Errorf("invalid date in CSV: %q, err: %w", record[12], err)
		}

		row := []interface{}{
			record[0], transactionDate, record[2], record[3], record[4],
			record[5], record[6], record[7], price, quantity, totalPrice,
			stockQuantity, addedDate,
		}

		batch = append(batch, row)

		if len(batch) >= 100_000 {
			if err := copyToStaging(conn, batch); err != nil {
				return err
			}
			batch = batch[:0]
		}
	}

	if len(batch) > 0 {
		if err := copyToStaging(conn, batch); err != nil {
			return err
		}
	}

	log.Println("CSV loaded into staging successfully.")

	// Normalize
	if err := normalizeData(conn); err != nil {
		return err
	}

	return nil
}

// copyToStaging streams a batch of rows into staging table
func copyToStaging(conn *pgx.Conn, rows [][]interface{}) error {
	_, err := conn.CopyFrom(
		context.Background(),
		pgx.Identifier{"staging"},
		[]string{"transaction_id", "transaction_date", "user_id", "country", "region",
			"product_id", "product_name", "category", "price", "quantity", "total_price",
			"stock_quantity", "added_date"},
		pgx.CopyFromRows(rows),
	)
	return err
}

// normalizeData moves data from staging to normalized tables
func normalizeData(conn *pgx.Conn) error {
	ctx := context.Background()

	queries := []string{
		// --- USERS ---
		// Pick the most recent record per user_id based on transaction_date
		`INSERT INTO users (user_id, country, region)
		 SELECT DISTINCT ON (user_id) user_id, country, region
		 FROM staging
		 ORDER BY user_id, transaction_date DESC
		 ON CONFLICT (user_id) DO UPDATE
		 SET country = EXCLUDED.country,
		     region = EXCLUDED.region;`,

		// --- PRODUCTS ---
		// Pick the most recent record per product_id based on added_date
		`INSERT INTO products (product_id, product_name, category, stock_quantity, added_date)
		 SELECT DISTINCT ON (product_id) product_id, product_name, category, stock_quantity, added_date
		 FROM staging
		 ORDER BY product_id, added_date DESC
		 ON CONFLICT (product_id) DO UPDATE
		 SET product_name = EXCLUDED.product_name,
		     category = EXCLUDED.category,
		     stock_quantity = EXCLUDED.stock_quantity,
		     added_date = EXCLUDED.added_date;`,

		// --- TRANSACTIONS ---
		// Insert all transactions (no updates, they are unique)
		`INSERT INTO transactions (transaction_id, transaction_date, user_id, product_id, price, quantity, total_price)
		 SELECT transaction_id, transaction_date, user_id, product_id, price, quantity, total_price
		 FROM staging
		 ON CONFLICT (transaction_id) DO NOTHING;`,
	}

	for _, q := range queries {
		if _, err := conn.Exec(ctx, q); err != nil {
			return fmt.Errorf("normalization failed: %v", err)
		}
	}

	log.Println("Data normalized successfully and correctly.")
	return nil
}
