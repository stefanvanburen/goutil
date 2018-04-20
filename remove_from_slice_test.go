package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveStringFromSlice(t *testing.T) {
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
		out, check := RemoveStringFromSlice(tc.find, tc.in)
		assert.Equal(t, out, tc.out, "Should be equal")
		assert.Equal(t, check, tc.expected, "Should be equal")
	}
}
