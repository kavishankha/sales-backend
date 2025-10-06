package dashboardhandler

import (
	"echo-app/internal/middlewares"
	"echo-app/internal/services/dashboardservice"
	"github.com/labstack/echo/v4"
	"net/http"
)

// RegionsByRevenueHandler Simple handler that calls the service function
func RegionsByRevenueHandler(c echo.Context) error {

	data, err := dashboardservice.RegionsByRevenue(c.Request().Context())
	if err != nil {
		return middlewares.Error(c, http.StatusInternalServerError, "Failed to fetch country revenue", err.Error())
	}

	return middlewares.Success(c, "Country revenue fetched successfully", data)
}
