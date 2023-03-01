package goutil

import (
	"testing"
)

func TestEditDistance(t *testing.T) {
	t.Parallel()
	tests := []struct {
		a    string
		b    string
		want int
	}{
		{"hello", "hello", 0},
		{"hello", "hell0", 1},
		{"hello", "hllo", 1},
		{"hello", "h3ll0", 2},
	}
	for _, test := range tests {
		test := test
		t.Run("EditDistance1", func(t *testing.T) {
			t.Parallel()
			if got := EditDistance1(test.a, test.b); got != test.want {
				t.Errorf("EditDistance1(%v, %v) = %v, want %v", test.a, test.b, got, test.want)
			}
		})
		t.Run("EditDistance2", func(t *testing.T) {
			t.Parallel()
			if got := EditDistance2(test.a, test.b); got != test.want {
				t.Errorf("EditDistance2(%v, %v) = %v, want %v", test.a, test.b, got, test.want)
			}
		})
	}
}
