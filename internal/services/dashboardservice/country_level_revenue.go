package dashboardservice

import (
	"context"
	"echo-app/internal/models"
	"echo-app/internal/repositories/dashboardrepo"
)

// CountryLevelRevenue Exported function: must start with uppercase C
func CountryLevelRevenue(ctx context.Context) ([]models.CountryReport, error) {
	return dashboardrepo.CountryLevelRevenue(ctx)
}
