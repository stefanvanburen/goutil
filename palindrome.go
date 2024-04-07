package goutil

// IsPalindrome checks if s is a palindrome.
func IsPalindrome(s string) bool {
	l := len(s)
	if l == 0 {
		return false
	}
	for i := range l / 2 {
		if s[i] != s[l-(i+1)] {
			return false
		}
	}
	return true
}
