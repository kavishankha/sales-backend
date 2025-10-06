package main

import (
	"context"
	"echo-app/config"
	"echo-app/pkg/database"
	"echo-app/routes"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// Start pprof server in a separate goroutine
	go func() {
		log.Println("Starting pprof on :6060")
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			log.Fatal("pprof server failed:", err)
		}
	}()

	config.LoadConfig()

	// Initialize app DB
	db := database.InitDB()
	defer db.Close()

	// 1️⃣ Run migrations
	database.RunMigrations(db)

	// 2️⃣ Create materialized views
	database.RunMaterializedViews(db)

	// 3️⃣ Load CSV into staging and normalize
	csvConn := database.GetPGXConn()
	defer csvConn.Close(context.Background())

	if err := database.LoadCSVAndSync(csvConn, os.Getenv("CSV_PATH")); err != nil {
		log.Fatal("Failed to load CSV:", err)
	}

	// 4️⃣ Refresh materialized views
	database.RefreshMaterializedViews(db)

	// Initialize Echo
	e := echo.New()

	// 🔐 Global panic recovery
	e.Use(middleware.Recover())

	// 🪵 Logging middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${uri} (${latency_human})\n",
	}))

	// 🌐 CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://192.168.1.101:3000", "http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// 📦 Register routes
	routes.RegisterRoutes(e)

	// 🚀 Start server
	port := config.GetEnv("PORT", "4000")
	log.Fatal(e.Start(":" + port))
}
