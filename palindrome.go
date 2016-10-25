package goutil

// IsPalindrome checks if a string is a palindrome
func IsPalindrome(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-(i+1)] {
			return false
		}
	}
	return true
}
