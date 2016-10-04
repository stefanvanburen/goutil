package goutil

import "testing"

func TestEditDistance(t *testing.T) {
	var tests = []struct {
		a   string
		b   string
		out int
	}{
		{"hello", "hello", 0},
		{"hello", "hell0", 1},
		{"hello", "hllo", 1},
		{"hello", "h3ll0", 2},
	}

	for _, x := range tests {
		c := EditDistance1(x.a, x.b)
		d := EditDistance2(x.a, x.b)
		if c != x.out {
			t.Errorf("EditDistance1(%s, %s) => %d, want %d", x.a, x.b, c, x.out)
		}
		if d != x.out {
			t.Errorf("EditDistance2(%s, %s) => %d, want %d", x.a, x.b, d, x.out)
		}
	}
}
