func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if s == goal {
		return true
	}
	return strings.Contains(s+s, goal)
}

func rotateString1(s string, goal string) bool {
	for i := 0; i <= len(s)-1; i++ {
		if s[i:]+s[:i] == goal {
			return true
		}
	}
	return false
}