package dashboardrepo

import (
	"context"
	"echo-app/internal/models"
	"echo-app/pkg/database"
	"fmt"

	"github.com/lib/pq"
)

func CountryLevelRevenue(ctx context.Context) ([]models.CountryReport, error) {
	query := `
		SELECT 
			country,
			products,
			total_revenue,
			transactions
		FROM country_report_mv
		ORDER BY total_revenue DESC;
	`

	rows, err := database.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var reports []models.CountryReport

	for rows.Next() {
		var r models.CountryReport
		var products pq.StringArray // scan Postgres text[] into pq.StringArray

		if err := rows.Scan(&r.Country, &products, &r.TotalRevenue, &r.Transactions); err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		println(products)
		r.Products = []string(products) // convert pq.StringArray to []string
		reports = append(reports, r)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return reports, nil
}
