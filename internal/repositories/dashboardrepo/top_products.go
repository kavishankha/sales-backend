package dashboardrepo

import (
	"context"
	"echo-app/internal/models"
	"echo-app/pkg/database"
)

func TopProduct(ctx context.Context) ([]models.Product, error) {
	query := `
		SELECT product_name, items_sold, stock_quantity
		FROM top_products_mv;
	`

	rows, err := database.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ProductName, &p.ItemsSold, &p.StockQuantity); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
