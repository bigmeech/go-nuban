package go_nuban

import (
	"testing"
)

// test validate function

func TestValidate(t *testing.T) {
	tests := []struct {
		name           string
		account_number string
		expected       bool
	}{
		{
			name:           "valid account number",
			account_number: "0000000000000",
			expected:       false,
		},
		{
			name:           "invalid account number",
			account_number: "00000000000000",
			expected:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Validate(tt.account_number)
			if actual != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}
