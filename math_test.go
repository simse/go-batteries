package batteries

import (
	"fmt"
	"testing"
)

func TestAbsInt(t *testing.T) {
	// Defining the columns of the table
	var tests = []struct {
		// name  string
		input int
		want  int
	}{
		// the table itself
		{9, 9},
		{-3, 3},
		{-1, 1},
		{0, 0},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d = %d", tt.input, tt.want), func(t *testing.T) {
			ans := AbsInt(tt.input)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
