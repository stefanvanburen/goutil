package goutil

// Divisors returns an integer slice of the proper divisors of an integer
func Divisors(n int) (divisors []int) {
	divisors = append(divisors, 1)
	for i := 2; i < SqrtInt(n); i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			divisors = append(divisors, n/i)
		}
	}
	return
}

// Abundant checks if an integer is an abundant integer
// An integer is abundant if the sum of its proper divisors is more than the integer
func Abundant(n int) bool {
	return Sum(Divisors(n)) > n
}

// Perfect checks if an integer is a perfect integer
// An integer is perfect if the sum of its proper divisors is equal to the integer
func Perfect(n int) bool {
	return Sum(Divisors(n)) == n
}

// Deficient checks if an integer is a deficient integer
// An integer is deficient if the sum of its proper divisors is less than the integer
func Deficient(n int) bool {
	return Sum(Divisors(n)) < n
}
