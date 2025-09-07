package main

import (
	"echo-app/config"
	"echo-app/pkg/database"
	"echo-app/routes"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	config.LoadConfig()
	// Initialize DB connection
	db := database.InitDB()
	if db == nil {
		log.Fatal("Failed to initialize database")
	}
	defer db.Close()

	// Run migrations
	database.RunMigrations(db)
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
