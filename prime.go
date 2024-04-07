package goutil

import "golang.org/x/exp/constraints"

// IsPrime checks if n is prime.
func IsPrime[K constraints.Integer](n K) bool {
	if n == 1 {
		return false
	}
	for b := SqrtInt(n); b >= 1; b-- {
		if b == 1 {
			return true
		}
		if n%b == 0 {
			return false
		}
	}
	return true
}

// PrimeSieve implements the Sieve of Eratosthenes.
func PrimeSieve[K constraints.Integer](max K) []K {
	bools := make([]bool, max)

	var primes []K

	var i K
	for ; i < max; i++ {
		bools[i] = true
	}

	// Sieve
	i = 2
	for ; i < max; i++ {
		if bools[i] {
			primes = append(primes, i)
			var k K = 2
			for ; i*k < max; k++ {
				bools[i*k] = false
			}
		}
	}

	return primes
}
