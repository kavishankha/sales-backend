package dashboardhandler

import (
	"context"
	"echo-app/internal/middlewares"
	"echo-app/internal/services/dashboardservice"
	"github.com/labstack/echo/v4"
	"time"
)

func CountryLevelRevenueHandler(c echo.Context) error {

	fetchFunc := func(ctx context.Context) (interface{}, error) {
		return dashboardservice.CountryLevelRevenue(ctx)
	}

	return middlewares.CacheMiddleware(
		"country_level_revenue",
		5*time.Minute,
		fetchFunc,
	)(c)
}
