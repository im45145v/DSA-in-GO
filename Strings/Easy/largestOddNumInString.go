func largestOddNumber(num string) string {
	var i int
	for i = len(num) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(num[i]))
		if digit%2 != 0 {
			return num[:i+1]
		}
	}
	return ""
}