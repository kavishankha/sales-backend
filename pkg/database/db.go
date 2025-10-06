package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

var (
	// Standard SQL DB for application queries
	DB *sql.DB
)

// GetDSN constructs the PostgreSQL connection string
func GetDSN() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, pass, host, port, name)
}

// InitDB initializes the standard *sql.DB connection for application use
func InitDB() *sql.DB {
	var err error
	DB, err = sql.Open("postgres", GetDSN())
	if err != nil {
		log.Fatalf("Error opening SQL DB: %v", err)
	}

	// Connection pool settings
	DB.SetMaxOpenConns(50)
	DB.SetMaxIdleConns(20)
	DB.SetConnMaxLifetime(time.Hour)

	if err := DB.Ping(); err != nil {
		log.Fatalf("SQL DB not reachable: %v", err)
	}

	log.Println("SQL DB connected successfully!")
	return DB
}

// GetPGXConn creates a high-performance pgx.Conn for bulk operations (CSV, normalization)
func GetPGXConn() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), GetDSN())
	if err != nil {
		log.Fatalf("Failed to connect PGX: %v", err)
	}
	log.Println("PGX connection established for high-performance operations")
	return conn
}
