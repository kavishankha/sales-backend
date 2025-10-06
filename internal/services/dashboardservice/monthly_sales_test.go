package dashboardservice

import (
	"testing"

	"echo-app/internal/models"
)

// TestCheckMonthLength tests the checkMonthLength function
func TestCheckMonthLength(t *testing.T) {
	tests := []struct {
		name      string
		count     int
		wantError bool
	}{
		{"ValidCount", 12, false},
		{"TooFewMonths", 10, true},
		{"TooManyMonths", 13, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkMonthLength(tt.count)
			if (err != nil) != tt.wantError {
				t.Errorf("checkMonthLength(%d) error = %v, wantError %v", tt.count, err, tt.wantError)
			}
		})
	}
}

// TestCheckDuplicateMonths tests the checkDuplicateMonths function
func TestCheckDuplicateMonths(t *testing.T) {
	tests := []struct {
		name      string
		sales     []models.MonthlySale
		wantError bool
	}{
		{
			"NoDuplicates",
			[]models.MonthlySale{
				{Month: "January"},
				{Month: "February"},
			},
			false,
		},
		{
			"WithDuplicates",
			[]models.MonthlySale{
				{Month: "January"},
				{Month: "January"},
			},
			true,
		},
		{
			"EmptySlice",
			[]models.MonthlySale{},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkDuplicateMonths(tt.sales)
			if (err != nil) != tt.wantError {
				t.Errorf("checkDuplicateMonths(%v) error = %v, wantError %v", tt.sales, err, tt.wantError)
			}
		})
	}
}
