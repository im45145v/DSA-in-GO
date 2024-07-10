func removeOuterParentheses(s string) string {
	count := 0
	result := make([]rune, len(s))
	for i, v := range s {
		if v == '(' {
			if count > 0 {
				result[i] = v
			}
			count++
		} else if v == ')' {
			count--
			if count > 0 {
				result[i] = v
			}
		}
	}
	return strings.ReplaceAll(string(result), "\x00", "")
}