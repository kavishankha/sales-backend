package dashboardservice

import (
	"context"
	"echo-app/internal/models"
	"echo-app/internal/repositories/dashboardrepo"
	"fmt"
)

// checkCountryLength ensures the number of countries does not exceed 195
func checkCountryLength(count int) error {
	if count > 195 {
		return fmt.Errorf("invalid data: more than 195 countries returned (%d)", count)
	}
	return nil
}

// checkDuplicateCountries ensures there are no duplicate country names
func checkDuplicateCountries(reports []models.CountryReport) error {
	countrySet := make(map[string]struct{})
	for _, r := range reports {
		if _, exists := countrySet[r.Country]; exists {
			return fmt.Errorf("invalid data: duplicate country found (%s)", r.Country)
		}
		countrySet[r.Country] = struct{}{}
	}
	return nil
}

// CountryLevelRevenue fetches country revenue and validates the reports
func CountryLevelRevenue(ctx context.Context) ([]models.CountryReport, error) {
	// First, fetch the reports from the repository
	reports, err := dashboardrepo.CountryLevelRevenue(ctx)
	if err != nil {
		return nil, err
	}

	// Validate length
	if err := checkCountryLength(len(reports)); err != nil {
		return nil, err
	}

	// Validate duplicates
	if err := checkDuplicateCountries(reports); err != nil {
		return nil, err
	}

	return reports, nil
}
