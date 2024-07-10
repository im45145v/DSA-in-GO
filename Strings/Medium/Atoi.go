
func myAtoi(s string) int {
	n, i, sign := len(s), 0, 1
	for i < n && s[i] == ' ' {
		i++
	}
	if i < n && s[i] == '-' {
		sign = -1
		i++
	} else if i < n && s[i] == '+' {
		i++
	}
	result := recur(s, i, sign, n, 0)
	return result
}

func recur(s string, i, sign, n int, ans int64) int {
	if ans*int64(sign) < int64(math.MinInt32) {
		return math.MinInt32
	}
	if ans*int64(sign) > int64(math.MaxInt32) {
		return math.MaxInt32
	}
	if i >= n || s[i] < '0' || s[i] > '9' {
		return int(ans) * sign
	}
	ans = ans*10 + int64(s[i]-'0')
	return recur(s, i+1, sign, n, ans)
}
