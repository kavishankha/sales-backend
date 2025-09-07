package dashboardrepo

import (
	"context"
	"echo-app/internal/models"
	"echo-app/pkg/database"
	"encoding/json"
)

func MonthlySales(ctx context.Context) ([]models.MonthlySale, error) {
	query := `SELECT jsonb_build_object(
    'monthlySales', jsonb_agg(month_data)
) AS result
FROM (
    SELECT 
        TO_CHAR(t.transaction_date, 'Mon') AS month,
        SUM(td.quantity) AS unitsSold
    FROM transaction_details td
    JOIN transactions t ON td.transaction_id = t.transaction_id
    GROUP BY TO_CHAR(t.transaction_date, 'Mon'), EXTRACT(MONTH FROM t.transaction_date)
    ORDER BY EXTRACT(MONTH FROM t.transaction_date)
) AS month_data;
`

	var rawJSON []byte
	if err := database.DB.QueryRowContext(ctx, query).Scan(&rawJSON); err != nil {
		return nil, err
	}

	var result struct {
		Data []models.MonthlySale `json:"monthlySales"`
	}
	if err := json.Unmarshal(rawJSON, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}
