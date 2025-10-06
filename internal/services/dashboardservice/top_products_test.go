package dashboardservice

import (
	"testing"

	"echo-app/internal/models"
)

// TestCheckProductLength tests the checkProductLength function
func TestCheckProductLength(t *testing.T) {
	tests := []struct {
		name      string
		count     int
		wantError bool
	}{
		{"ValidCount", 10, false},
		{"ExactLimit", 20, false},
		{"ExceedLimit", 21, true},
		{"ZeroProducts", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkProductLength(tt.count)
			if (err != nil) != tt.wantError {
				t.Errorf("checkProductLength(%d) error = %v, wantError %v", tt.count, err, tt.wantError)
			}
		})
	}
}

// TestCheckDuplicateProducts tests the checkDuplicateProducts function
func TestCheckDuplicateProducts(t *testing.T) {
	tests := []struct {
		name      string
		products  []models.Product
		wantError bool
	}{
		{
			"NoDuplicates",
			[]models.Product{
				{ProductName: "ProductA"},
				{ProductName: "ProductB"},
			},
			false,
		},
		{
			"WithDuplicates",
			[]models.Product{
				{ProductName: "ProductA"},
				{ProductName: "ProductA"},
			},
			true,
		},
		{
			"EmptySlice",
			[]models.Product{},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkDuplicateProducts(tt.products)
			if (err != nil) != tt.wantError {
				t.Errorf("checkDuplicateProducts(%v) error = %v, wantError %v", tt.products, err, tt.wantError)
			}
		})
	}
}
