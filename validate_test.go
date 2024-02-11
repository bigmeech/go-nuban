package go_nuban

import (
	"testing"
)

// test validate function

func TestValidate(t *testing.T) {
	tests := []struct {
		name  string
		nuban string
		want  bool
	}{
		{
			name:  "Valid",
			nuban: "0000014579",
			want:  true,
		},
		{
			name:  "Invalid",
			nuban: "0000014578",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Validate(tt.nuban); got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
