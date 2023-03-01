package goutil

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestRemoveFromSlice(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in          []string
		find        string
		wantSlice   []string
		wantRemoved bool
	}{
		{[]string{"hey", "bye", "hi"}, "hi", []string{"hey", "bye"}, true},
		{[]string{"hey", "bye", "hi"}, "not in", []string{"hey", "bye", "hi"}, false},
	}
	for _, tc := range tests {
		gotSlice, gotRemoved := RemoveFromSlice(tc.find, tc.in)
		if !slices.Equal(gotSlice, tc.wantSlice) {
			t.Errorf("RemoveFromSlice(%v) = %v, want %v", tc.find, gotSlice, tc.wantSlice)
		}
		if gotRemoved != tc.wantRemoved {
			t.Errorf("RemoveFromSlice(%v) = %v, want %v", tc.find, gotRemoved, tc.wantRemoved)
		}
	}
}
