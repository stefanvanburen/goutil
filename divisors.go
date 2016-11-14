package goutil

// Divisors returns an integer slice of the proper divisors of an integer
func Divisors(n int) (divisors []int) {
	count := 1
	divisors = append(divisors, 1)
	for i := 2; i <= SqrtInt(n); i++ {
		if n%i == 0 {
			// Rather than sorting at the end, we can take advantage of the
			// fact that i > previous i values and n/i < previous n/i values
			// to keep the divisors array sorted
			// So, we insert i and n/i in consecutive positions after the previous
			// i value's location

			// Free up space for two more divisors
			divisors = append(divisors, 0)
			divisors = append(divisors, 0)

			// Copy the values past count back two positions
			copy(divisors[count+2:], divisors[count:])

			divisors[count] = i
			divisors[count+1] = n / i
			count++
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
