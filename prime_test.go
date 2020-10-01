package goutil

import (
	"testing"

	"github.com/matryer/is"
)

func TestPrimeSieve(t *testing.T) {
	is := is.New(t)

	x := PrimeSieve(1000)
	for _, e := range x {
		is.True(IsPrime(e))
	}
}
