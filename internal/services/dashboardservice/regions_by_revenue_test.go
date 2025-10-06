package dashboardservice

import (
	"testing"
)

// TestCheckRegionLength tests the checkRegionLength function
func TestCheckRegionLength(t *testing.T) {
	tests := []struct {
		name      string
		count     int
		wantError bool
	}{
		{"ValidCount", 10, false},
		{"ExactLimit", 30, false},
		{"ExceedLimit", 31, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkRegionLength(tt.count)
			if (err != nil) != tt.wantError {
				t.Errorf("checkRegionLength(%d) error = %v, wantError %v", tt.count, err, tt.wantError)
			}
		})
	}
}
