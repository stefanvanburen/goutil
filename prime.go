package goutil

// IsPrime checks if an int is prime
func IsPrime(n int) bool {
	for b := Sqrt(n); b >= 1; b-- {
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
	for b := Sqrt64(n); b >= 1; b-- {
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
func PrimeSieve(len int) []int {
	bools := make([]bool, len)
	var primes []int
	var i, count int
	for i = 0; i < len; i++ {
		bools[i] = true
	}
	// Sieve
	for i = 2; i < len; i++ {
		if bools[i] {
			count++
			primes = append(primes, i)
			for k := 2; i*k < len; k++ {
				bools[i*k] = false
			}
		}
	}
	return primes
}
