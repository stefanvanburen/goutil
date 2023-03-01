package goutil

import (
	"testing"
)

func TestPrimeSieve(t *testing.T) {
	t.Parallel()
	for _, got := range PrimeSieve(1000) {
		if !IsPrime(got) {
			t.Errorf("IsPrime(%v) from PrimeSieve(1000) = false", got)
		}
	}
}
