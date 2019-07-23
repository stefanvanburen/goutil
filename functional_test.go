package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	var tests = map[string]struct {
		haystack []string
		needle   string
		expected int
	}{
		"exists":           {[]string{"one", "two"}, "two", 1},
		"doesn't exist":    {[]string{"one", "two"}, "zero", -1},
		"multiple matches": {[]string{"one", "two", "two"}, "two", 1},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, Index(tc.haystack, tc.needle), tc.expected)
		})
	}
}

func TestInclude(t *testing.T) {
	var tests = map[string]struct {
		haystack []string
		needle   string
		expected bool
	}{
		"exists":           {[]string{"one", "two"}, "two", true},
		"doesn't exist":    {[]string{"one", "two"}, "zero", false},
		"multiple matches": {[]string{"one", "two", "two"}, "two", true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, Include(tc.haystack, tc.needle), tc.expected)
		})
	}
}
