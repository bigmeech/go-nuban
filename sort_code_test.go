package go_nuban

import "testing"

// test generateSortCode function using fiCode, state code and branch code
func TestGenerateSortCode(t *testing.T) {
	tests := []struct {
		name       string
		fiCode     string
		stateCode  string
		branchCode string
		want       string
	}{
		{
			name:       "DMB",
			fiCode:     "011",
			stateCode:  "01",
			branchCode: "457",
			want:       "011-01-457-7",
		},
		{
			name:       "OFI",
			fiCode:     "50547",
			stateCode:  "50",
			branchCode: "457",
			want:       "950547-50-457-3",
		},
		{
			name:       "Invalid",
			fiCode:     "1234",
			stateCode:  "12",
			branchCode: "345",
			want:       "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateSortCode(tt.fiCode, tt.stateCode, tt.branchCode); got != tt.want {
				t.Errorf("GenerateSortCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
