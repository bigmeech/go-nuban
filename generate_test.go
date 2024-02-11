package go_nuban

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name         string
		fiCode       string
		serialNumber string
		want         string
	}{
		{
			name:         "DMB",
			fiCode:       "011",
			serialNumber: "000001457",
			want:         "0000014579",
		},
		{
			name:         "OFI",
			fiCode:       "50547",
			serialNumber: "000021457",
			want:         "0000214579",
		},
		{
			name:         "Invalid",
			fiCode:       "1234",
			serialNumber: "123456",
			want:         "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Generate(tt.fiCode, tt.serialNumber); got != tt.want {
				t.Errorf("Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
