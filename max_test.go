package goutil

import (
	"testing"
)

func TestMaxInt(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   []int
		want int
	}{
		{[]int{-17, 5, 2, 18, 20, -64, 8}, 20},
		{[]int{-10, -9, -8, -6}, -6},
	}
	for _, tt := range tests {
		if got := Max(tt.in...); got != tt.want {
			t.Errorf("Max(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}
}
