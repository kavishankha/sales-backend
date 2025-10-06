package dashboardrepo

import (
	"context"
	"echo-app/internal/models"
	"echo-app/pkg/database"
)

func RegionsByRevenue(ctx context.Context) ([]models.Region, error) {
	query := `
		SELECT region, revenue_usd, items_sold
		FROM regions_by_revenue_mv;
	`

	rows, err := database.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var regions []models.Region
	for rows.Next() {
		var r models.Region
		if err := rows.Scan(&r.Region, &r.RevenueUSD, &r.ItemsSold); err != nil {
			return nil, err
		}
		regions = append(regions, r)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return regions, nil
}
