package dashboardservice

import (
	"context"
	"echo-app/internal/models"
	"echo-app/internal/repositories/dashboardrepo"
)

// TopProduct Exported function: must start with uppercase C
func TopProduct(ctx context.Context) ([]models.Product, error) {
	return dashboardrepo.TopProduct(ctx)
}
