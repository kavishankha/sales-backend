package dashboardservice

import (
	"testing"

	"echo-app/internal/models"
)

// TestCheckCountryLength tests the checkCountryLength function
func TestCheckCountryLength(t *testing.T) {
	tests := []struct {
		name      string
		count     int
		wantError bool
	}{
		{"ValidCount", 50, false},
		{"ExactLimit", 195, false},
		{"ExceedLimit", 196, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkCountryLength(tt.count)
			if (err != nil) != tt.wantError {
				t.Errorf("checkCountryLength(%d) error = %v, wantError %v", tt.count, err, tt.wantError)
			}
		})
	}
}

// TestCheckDuplicateCountries tests the checkDuplicateCountries function
func TestCheckDuplicateCountries(t *testing.T) {
	tests := []struct {
		name      string
		reports   []models.CountryReport
		wantError bool
	}{
		{
			"NoDuplicates",
			[]models.CountryReport{
				{Country: "USA"},
				{Country: "China"},
			},
			false,
		},
		{
			"WithDuplicates",
			[]models.CountryReport{
				{Country: "USA"},
				{Country: "USA"},
			},
			true,
		},
		{
			"EmptySlice",
			[]models.CountryReport{},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkDuplicateCountries(tt.reports)
			if (err != nil) != tt.wantError {
				t.Errorf("checkDuplicateCountries(%v) error = %v, wantError %v", tt.reports, err, tt.wantError)
			}
		})
	}
}
