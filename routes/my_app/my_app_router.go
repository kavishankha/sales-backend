package my_app

import (
	"github.com/labstack/echo/v4"
)

func MyAppRoutes(e *echo.Echo) {
	MyApp := e.Group("/my_app")
	DashboardRoutes(MyApp)
	//AuthRoutes(MyApp)
}
