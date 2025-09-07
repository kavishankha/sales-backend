package dashboardrepo

import (
	"context"
	"echo-app/internal/models"
	"echo-app/pkg/database"
	"encoding/json"
)

func RegionsByRevenue(ctx context.Context) ([]models.Region, error) {
	query := `SELECT jsonb_build_object(
    'data', jsonb_agg(region_data)
) AS result
FROM (
    SELECT 
        u.region,
        SUM(td.total_price) AS revenueUSD,
        SUM(td.quantity) AS itemsSold
    FROM transaction_details td
    JOIN transactions t ON td.transaction_id = t.transaction_id
    JOIN users u ON t.user_id = u.user_id
    GROUP BY u.region
    ORDER BY SUM(td.total_price) DESC
    LIMIT 30
) AS region_data;
`

	var rawJSON []byte
	if err := database.DB.QueryRowContext(ctx, query).Scan(&rawJSON); err != nil {
		return nil, err
	}

	var result struct {
		Data []models.Region `json:"data"`
	}
	if err := json.Unmarshal(rawJSON, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}
