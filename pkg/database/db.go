package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// GetDSN returns PostgreSQL connection string
func GetDSN() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, pass, host, port, name)
}

// InitDB initializes standard *sql.DB for app queries
func InitDB() *sql.DB {
	var err error
	DB, err = sql.Open("postgres", GetDSN())
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}

	// Connection pool
	DB.SetMaxOpenConns(50)
	DB.SetMaxIdleConns(20)
	DB.SetConnMaxLifetime(time.Hour)

	if err := DB.Ping(); err != nil {
		log.Fatalf("DB not reachable: %v", err)
	}

	log.Println("Database connected successfully!")
	return DB
}

// GetPGXConn returns a pgx.Conn for high-performance operations
func GetPGXConn() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), GetDSN())
	if err != nil {
		log.Fatalf("Failed to connect pgx: %v", err)
	}
	return conn
}
