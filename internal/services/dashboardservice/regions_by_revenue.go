package dashboardservice

import (
	"context"
	"echo-app/internal/models"
	"echo-app/internal/repositories/dashboardrepo"
	"fmt"
)

// checkRegionLength ensures the number of regions does not exceed 30
func checkRegionLength(count int) error {
	if count > 30 {
		return fmt.Errorf("invalid data: more than 30 regions returned (%d)", count)
	}
	return nil
}

// RegionsByRevenue fetches regions by revenue and validates the number of regions
func RegionsByRevenue(ctx context.Context) ([]models.Region, error) {
	// Fetch data from repository
	regions, err := dashboardrepo.RegionsByRevenue(ctx)
	if err != nil {
		return nil, err
	}

	// Check length
	if err := checkRegionLength(len(regions)); err != nil {
		return nil, err
	}

	return regions, nil
}
