func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	res := fmt.Sprintf("%c", s[0])
	isPal := func(str string) bool {
		if len(str) < 2 {
			return true
		}
		l, r := 0, len(str)-1
		for ; l < r; l, r = l+1, r-1 {
			if str[l] != str[r] {
				return false
			}
		}
		return true
	}
	m := make(map[byte][]int)
	for i := 0; i < len(s); i++ {
		m[s[i]] = append(m[s[i]], i)
		if len(m[s[i]]) > 1 {
			for _, val := range m[s[i]] {
				if i != val && isPal(s[val:i+1]) {
					if i-val+1 > len(res) {
						res = s[val : i+1]
					}
				}
			}
		}
	}
	return res
}