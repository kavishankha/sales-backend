package dashboardhandler

import (
	"echo-app/internal/middlewares"
	"echo-app/internal/services/dashboardservice"
	"github.com/labstack/echo/v4"
	"net/http"
)

// MonthlySalesHandler Simple handler that calls the service function
func MonthlySalesHandler(c echo.Context) error {

	data, err := dashboardservice.MonthlySales(c.Request().Context())
	if err != nil {
		return middlewares.Error(c, http.StatusInternalServerError, "Failed to fetch country revenue", err.Error())
	}

	return middlewares.Success(c, "Country revenue fetched successfully", data)
}
