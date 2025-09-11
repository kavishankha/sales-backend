package dashboardrepo

import (
	"context"
	"echo-app/internal/models"
	"echo-app/pkg/database"
	"encoding/json"
)

func MonthlySales(ctx context.Context) ([]models.MonthlySale, error) {
	query := `SELECT jsonb_build_object(
    'monthlySales', jsonb_agg(month_data ORDER BY month_num)
) AS result
FROM (
    SELECT 
        TO_CHAR(t.transaction_date, 'Mon') AS month,
        DATE_PART('month', t.transaction_date) AS month_num,
        SUM(t.quantity) AS unitsSold
    FROM transactions t
    GROUP BY month, month_num
    ORDER BY month_num
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
