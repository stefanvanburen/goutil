package goutil

// Sum returns the sum of vals.
func Sum[K Integer](vals ...K) (sum K) {
	for _, v := range vals {
		sum += v
	}
	return sum
}
