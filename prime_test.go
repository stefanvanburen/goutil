package goutil

import "testing"

func TestPrimeSieve(t *testing.T) {
	x := PrimeSieve(1000)
	for _, e := range x {
		if !IsPrime(e) {
			t.Errorf("Not prime element %d", e)
		}
	}
}
