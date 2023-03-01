package goutil

import (
	"testing"
)

func TestMinInt(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   []int
		want int
	}{
		{[]int{-17, 5, 2, 18, 20, -64, 8}, -64},
		{[]int{-10, -9, -8, -6}, -10},
	}
	for _, tt := range tests {
		if got := Min(tt.in...); got != tt.want {
			t.Errorf("Min(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}
}
