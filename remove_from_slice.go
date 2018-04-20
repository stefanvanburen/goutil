package goutil

// RemoveStringFromSlice removes the first instance of a string from the given slice.
// It returns the slice and a bool to indicate if a value was removed or not.
func RemoveStringFromSlice(s string, sl []string) ([]string, bool) {
	for i, v := range sl {
		if v == s {
			return append(sl[:i], sl[i+1:]...), true
		}
	}
	return sl, false
}
