package goutil

// IsPrime checks if an int is prime
func IsPrime(n int) bool {
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

// IsPrime64 checks if an int64 is prime
func IsPrime64(n int64) bool {
	if n == 1 {
		return false
	}
	for b := SqrtInt64(n); b >= 1; b-- {
		if b == 1 {
			return true
		}
		if n%b == 0 {
			return false
		}
	}
	return true
}

// PrimeSieve implements the Sieve of Eratosthenes
func PrimeSieve(max int) []int {
	bools := make([]bool, max)
	var primes []int
	var i int
	for i = 0; i < max; i++ {
		bools[i] = true
	}
	// Sieve
	for i = 2; i < max; i++ {
		if bools[i] {
			primes = append(primes, i)
			for k := 2; i*k < max; k++ {
				bools[i*k] = false
			}
		}
	}
	return primes
}
