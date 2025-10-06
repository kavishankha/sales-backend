package dashboardservice

import (
	"context"
	"echo-app/internal/models"
	"echo-app/internal/repositories/dashboardrepo"
	"fmt"
)

// checkMonthLength ensures there are exactly 12 months
func checkMonthLength(count int) error {
	if count != 12 {
		return fmt.Errorf("invalid data: number of months must be 12, got %d", count)
	}
	return nil
}

// checkDuplicateMonths ensures there are no duplicate month names
func checkDuplicateMonths(sales []models.MonthlySale) error {
	monthSet := make(map[string]struct{})
	for _, s := range sales {
		if _, exists := monthSet[s.Month]; exists {
			return fmt.Errorf("invalid data: duplicate month found (%s)", s.Month)
		}
		monthSet[s.Month] = struct{}{}
	}
	return nil
}

// MonthlySales fetches monthly sales and validates length and duplicates
func MonthlySales(ctx context.Context) ([]models.MonthlySale, error) {
	// Fetch data from repository
	sales, err := dashboardrepo.MonthlySales(ctx)
	if err != nil {
		return nil, err
	}

	// Check length
	if err := checkMonthLength(len(sales)); err != nil {
		return nil, err
	}

	// Check duplicates
	if err := checkDuplicateMonths(sales); err != nil {
		return nil, err
	}

	return sales, nil
}
