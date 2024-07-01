package roman

import "testing"

func TestRomanToNumber(t *testing.T) {
	
	tests := []struct {
		name string
		args string
		want int
	}{
		{name: "case 1", args: "I", want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RomanToNumber(tt.args); got != tt.want {
				t.Errorf("RomanToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
