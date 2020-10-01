package goutil

import (
	"testing"

	"github.com/matryer/is"
)

func TestIndex(t *testing.T) {
	tests := map[string]struct {
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
			t.Parallel()
			is := is.New(t)

			is.Equal(Index(tc.haystack, tc.needle), tc.expected)
		})
	}
}

func TestInclude(t *testing.T) {
	tests := map[string]struct {
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
			t.Parallel()
			is := is.New(t)

			is.Equal(Include(tc.haystack, tc.needle), tc.expected)
		})
	}
}

func TestAny(t *testing.T) {
	hasLength := func(s string) bool {
		return len(s) > 0
	}

	tests := map[string]struct {
		haystack  []string
		predicate func(string) bool
		expected  bool
	}{
		"no matches": {[]string{"", ""}, hasLength, false},
		"one match":  {[]string{"hi", ""}, hasLength, true},
		"all match":  {[]string{"hi", "bye"}, hasLength, true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			is.Equal(Any(tc.haystack, tc.predicate), tc.expected)
		})
	}
}

func TestAll(t *testing.T) {
	hasLength := func(s string) bool {
		return len(s) > 0
	}

	tests := map[string]struct {
		haystack  []string
		predicate func(string) bool
		expected  bool
	}{
		"no matches": {[]string{"", ""}, hasLength, false},
		"one match":  {[]string{"hi", ""}, hasLength, false},
		"all match":  {[]string{"hi", "bye"}, hasLength, true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			is.Equal(All(tc.haystack, tc.predicate), tc.expected)
		})
	}
}

func TestFilter(t *testing.T) {
	hasLength := func(s string) bool {
		return len(s) > 0
	}

	tests := map[string]struct {
		haystack  []string
		predicate func(string) bool
		expected  []string
	}{
		"no matches": {[]string{"", ""}, hasLength, nil},
		"one match":  {[]string{"hi", ""}, hasLength, []string{"hi"}},
		"all match":  {[]string{"hi", "bye"}, hasLength, []string{"hi", "bye"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			is.Equal(Filter(tc.haystack, tc.predicate), tc.expected)
		})
	}
}

func TestMap(t *testing.T) {
	exclamation := func(s string) string {
		return s + "!"
	}

	tests := map[string]struct {
		haystack  []string
		transform func(string) string
		expected  []string
	}{
		"no matches": {[]string{"", ""}, exclamation, []string{"!", "!"}},
		"one match":  {[]string{"hi", ""}, exclamation, []string{"hi!", "!"}},
		"all match":  {[]string{"hi", "bye"}, exclamation, []string{"hi!", "bye!"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			is.Equal(Map(tc.haystack, tc.transform), tc.expected)
		})
	}
}
