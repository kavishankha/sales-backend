package dashboardrepo

import (
	"context"
	"echo-app/internal/models"
	"echo-app/pkg/database"
)

func MonthlySales(ctx context.Context) ([]models.MonthlySale, error) {
	query := `
		SELECT month, units_sold
		FROM monthly_sales_mv
		ORDER BY month_num;
	`

	rows, err := database.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sales []models.MonthlySale
	for rows.Next() {
		var s models.MonthlySale
		if err := rows.Scan(&s.Month, &s.UnitsSold); err != nil {
			return nil, err
		}
		sales = append(sales, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sales, nil
}
