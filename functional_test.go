package goutil

import (
	"slices"
	"testing"
)

func TestIndex(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		haystack []string
		needle   string
		want     int
	}{
		"exists":           {[]string{"one", "two"}, "two", 1},
		"doesn't exist":    {[]string{"one", "two"}, "zero", -1},
		"multiple matches": {[]string{"one", "two", "two"}, "two", 1},
	}
	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := Index(tc.haystack, tc.needle); got != tc.want {
				t.Errorf("Index(%v, %v) = %v, want %v", tc.haystack, tc.needle, got, tc.want)
			}
		})
	}
}

func TestInclude(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		haystack []string
		needle   string
		want     bool
	}{
		"exists":           {[]string{"one", "two"}, "two", true},
		"doesn't exist":    {[]string{"one", "two"}, "zero", false},
		"multiple matches": {[]string{"one", "two", "two"}, "two", true},
	}
	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := Include(tc.haystack, tc.needle); got != tc.want {
				t.Errorf("Include(%v, %v) = %v, want %v", tc.haystack, tc.needle, got, tc.want)
			}
		})
	}
}

func TestAny(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		haystack  []string
		predicate func(string) bool
		want      bool
	}{
		"no matches": {[]string{"", ""}, hasLength, false},
		"one match":  {[]string{"hi", ""}, hasLength, true},
		"all match":  {[]string{"hi", "bye"}, hasLength, true},
	}
	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := Any(tc.haystack, tc.predicate); got != tc.want {
				t.Errorf("Any(%v, %T) = %v, want %v", tc.haystack, tc.predicate, got, tc.want)
			}
		})
	}
}

func TestAll(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		haystack  []string
		predicate func(string) bool
		want      bool
	}{
		"no matches": {[]string{"", ""}, hasLength, false},
		"one match":  {[]string{"hi", ""}, hasLength, false},
		"all match":  {[]string{"hi", "bye"}, hasLength, true},
	}
	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := All(tc.haystack, tc.predicate); got != tc.want {
				t.Errorf("All(%v, %T) = %v, want %v", tc.haystack, tc.predicate, got, tc.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		haystack  []string
		predicate func(string) bool
		want      []string
	}{
		"no matches": {[]string{"", ""}, hasLength, nil},
		"one match":  {[]string{"hi", ""}, hasLength, []string{"hi"}},
		"all match":  {[]string{"hi", "bye"}, hasLength, []string{"hi", "bye"}},
	}
	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := Filter(tc.haystack, tc.predicate); !slices.Equal(got, tc.want) {
				t.Errorf("Filter(%v, %T) = %v, want %v", tc.haystack, tc.predicate, got, tc.want)
			}
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
		want      []string
	}{
		"no matches": {[]string{"", ""}, exclamation, []string{"!", "!"}},
		"one match":  {[]string{"hi", ""}, exclamation, []string{"hi!", "!"}},
		"all match":  {[]string{"hi", "bye"}, exclamation, []string{"hi!", "bye!"}},
	}
	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := Map(tc.haystack, tc.transform); !slices.Equal(got, tc.want) {
				t.Errorf("Map(%v, %T) = %v, want %v", tc.haystack, tc.transform, got, tc.want)
			}
		})
	}
}

func hasLength(s string) bool {
	return len(s) > 0
}
