package dashboardservice

import (
	"context"
	"echo-app/internal/models"
	"echo-app/internal/repositories/dashboardrepo"
)

// TopProduct Exported function: must start with uppercase C
func RegionsByRevenue(ctx context.Context) ([]models.Region, error) {
	return dashboardrepo.RegionsByRevenue(ctx)
}
