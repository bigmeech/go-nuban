package go_nuban

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name          string
		inst_code     int
		serial_number int
		expected      string
	}{
		{
			name:          "valid account number",
			inst_code:     0,
			serial_number: 0,
			expected:      "0000000000000",
		},
		{
			name:          "invalid account number",
			inst_code:     0,
			serial_number: 0,
			expected:      "00000000000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Generate(tt.inst_code, tt.serial_number)
			if actual != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}
