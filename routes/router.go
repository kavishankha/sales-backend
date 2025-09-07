package routes

import (
	"echo-app/routes/my_app"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	my_app.MyAppRoutes(e)

}
