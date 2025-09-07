package dashboardrepo

import (
	"context"
	"echo-app/internal/models"
	"echo-app/pkg/database"
	"encoding/json"
)

func CountryLevelRevenue(ctx context.Context) ([]models.CountryReport, error) {
	query := `
		SELECT jsonb_build_object(
			'data', jsonb_agg(country_data)
		) AS result
		FROM (
			SELECT 
				u.country,
				jsonb_agg(DISTINCT p.product_name) AS products,
				SUM(td.total_price) AS totalRevenue,
				COUNT(DISTINCT t.transaction_id) AS transactions
			FROM transaction_details td
			JOIN transactions t ON td.transaction_id = t.transaction_id
			JOIN users u ON t.user_id = u.user_id
			JOIN products p ON td.product_id = p.product_id
			GROUP BY u.country
			ORDER BY SUM(td.total_price) DESC
		) AS country_data;
	`

	var rawJSON []byte
	if err := database.DB.QueryRowContext(ctx, query).Scan(&rawJSON); err != nil {
		return nil, err
	}

	var result struct {
		Data []models.CountryReport `json:"data"`
	}
	if err := json.Unmarshal(rawJSON, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}
