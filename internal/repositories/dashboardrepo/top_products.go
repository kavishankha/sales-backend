package dashboardrepo

import (
	"context"
	"echo-app/internal/models"
	"echo-app/pkg/database"
	"encoding/json"
)

func TopProduct(ctx context.Context) ([]models.Product, error) {
	query := `
SELECT jsonb_build_object(
    'data', jsonb_agg(product_data)
) AS result
FROM (
    SELECT 
        p.product_name,
        SUM(t.quantity) AS itemsSold,
        p.stock_quantity
    FROM transactions t
    JOIN products p ON t.product_id = p.product_id
    GROUP BY p.product_id, p.product_name, p.stock_quantity
    ORDER BY itemsSold DESC
    LIMIT 20
) AS product_data;
`

	var rawJSON []byte
	if err := database.DB.QueryRowContext(ctx, query).Scan(&rawJSON); err != nil {
		return nil, err
	}

	var result struct {
		Data []models.Product `json:"data"`
	}
	if err := json.Unmarshal(rawJSON, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}
