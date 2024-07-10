func maxDepth(s string) int {
	ultmax, curr := 0, 0
	for _, v := range s {
		if v == '(' {
			curr += 1
			if curr > ultmax {
				ultmax = curr
			}
		} else if v == ')' {
			curr -= 1
		}
	}
	return ultmax
}