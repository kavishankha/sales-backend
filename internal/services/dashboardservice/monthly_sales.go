package dashboardservice

import (
	"context"
	"echo-app/internal/models"
	"echo-app/internal/repositories/dashboardrepo"
)

// TopProduct Exported function: must start with uppercase C
func MonthlySales(ctx context.Context) ([]models.MonthlySale, error) {
	return dashboardrepo.MonthlySales(ctx)
}
