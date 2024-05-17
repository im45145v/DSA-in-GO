func isPalindrome(s string) bool {
	for len(s) > 0 && !isalphanum(rune(s[0])) {
		s = s[1:]
	}
	for len(s) > 0 && !isalphanum(rune(s[len(s)-1])) {
		s = s[:len(s)-1]
	}
	if len(s) <= 1 {
		return true
	} else {
		lo := unicode.ToLower(rune(s[0]))
		hi := unicode.ToLower(rune(s[len(s)-1]))
		return lo == hi && isPalindrome(s[1:len(s)-1])
	}
}
func isalphanum(ch rune) bool {
	if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
		return true
	}
	return false
}