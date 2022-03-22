package goutil

// RemoveFromSlice removes the first instance of s from the given slice.
// It returns the (possibly modified) slice and a bool to indicate if a value
// was removed or not.
func RemoveFromSlice[K comparable](s K, sl []K) ([]K, bool) {
	for i, v := range sl {
		if v == s {
			return append(sl[:i], sl[i+1:]...), true
		}
	}

	return sl, false
}
