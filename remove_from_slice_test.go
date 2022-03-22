package goutil

import (
	"testing"

	"github.com/matryer/is"
)

func TestRemoveStringFromSlice(t *testing.T) {
	is := is.New(t)

	tcs := []struct {
		in       []string
		find     string
		out      []string
		expected bool
	}{
		{[]string{"hey", "bye", "hi"}, "hi", []string{"hey", "bye"}, true},
		{[]string{"hey", "bye", "hi"}, "not in", []string{"hey", "bye", "hi"}, false},
	}

	for _, tc := range tcs {
		out, check := RemoveFromSlice(tc.find, tc.in)
		is.Equal(out, tc.out)
		is.Equal(check, tc.expected)
	}
}
