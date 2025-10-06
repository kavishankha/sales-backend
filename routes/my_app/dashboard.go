package my_app

import (
	"echo-app/internal/handlers/dashboardhandler"

	"github.com/labstack/echo/v4"
)

// DashboardRoutes sets up the routes for the dashboard group.
func DashboardRoutes(MyAppGroup *echo.Group) {

	dashboard := MyAppGroup.Group("/dashboard")

	dashboard.GET("/countries", dashboardhandler.CountryLevelRevenueHandler)
	dashboard.GET("/top-products", dashboardhandler.TopProductHandler)
	dashboard.GET("/monthly-sales", dashboardhandler.MonthlySalesHandler)
	dashboard.GET("/regions", dashboardhandler.RegionsByRevenueHandler)
}
