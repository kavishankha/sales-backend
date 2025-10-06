package dashboardhandler

import (
	"echo-app/internal/middlewares"
	"echo-app/internal/services/dashboardservice"
	"github.com/labstack/echo/v4"
	"net/http"
)

//func CountryLevelRevenueHandler(c echo.Context) error {
//
//	fetchFunc := func(ctx context.Context) (interface{}, error) {
//		return dashboardservice.CountryLevelRevenue(ctx)
//	}
//
//	return middlewares.CacheMiddleware(
//		"country_level_revenue",
//		5*time.Minute,
//		fetchFunc,
//	)(c)
//}

// CountryLevelRevenueHandler Simple handler that calls the service function
func CountryLevelRevenueHandler(c echo.Context) error {

	data, err := dashboardservice.CountryLevelRevenue(c.Request().Context())
	if err != nil {
		return middlewares.Error(c, http.StatusInternalServerError, "Failed to fetch country revenue", err.Error())
	}

	return middlewares.Success(c, "Country revenue fetched successfully", data)
}
