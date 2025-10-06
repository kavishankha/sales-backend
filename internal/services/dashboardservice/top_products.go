package dashboardservice

import (
	"context"
	"echo-app/internal/models"
	"echo-app/internal/repositories/dashboardrepo"
	"fmt"
)

// checkProductLength ensures the number of products does not exceed 20
func checkProductLength(count int) error {
	if count > 20 {
		return fmt.Errorf("invalid data: more than 20 products returned (%d)", count)
	}
	return nil
}

// checkDuplicateProducts ensures there are no duplicate product names
func checkDuplicateProducts(products []models.Product) error {
	productSet := make(map[string]struct{})
	for _, p := range products {
		if _, exists := productSet[p.ProductName]; exists {
			return fmt.Errorf("invalid data: duplicate product found (%s)", p.ProductName)
		}
		productSet[p.ProductName] = struct{}{}
	}
	return nil
}

// TopProduct fetches top products and validates length and duplicates
func TopProduct(ctx context.Context) ([]models.Product, error) {
	// Fetch data from repository
	products, err := dashboardrepo.TopProduct(ctx)
	if err != nil {
		return nil, err
	}

	// Check length
	if err := checkProductLength(len(products)); err != nil {
		return nil, err
	}

	// Check duplicates
	if err := checkDuplicateProducts(products); err != nil {
		return nil, err
	}

	return products, nil
}
