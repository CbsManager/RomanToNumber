package roman

import "testing"

func TestRomanToNumber(t *testing.T) {
	
	tests := []struct {
		name string
		args string
		want int
	}{
		{name: "case 1", args: "I", want: 1},
		{name: "case 2", args: "II", want: 2},
		{name: "case 4", args: "IV", want: 4},
		{name: "case 9", args: "IX", want: 9},
		{name: "case 14", args: "XIV", want: 14},
		{name: "case 44", args: "XLIV", want: 44},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RomanToNumber(tt.args); got != tt.want {
				t.Errorf("RomanToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
