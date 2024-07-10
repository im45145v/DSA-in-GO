func romanToInt(s string) int {
	ourMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	val := 0
	for i, v := range s {
		val += ourMap[string(v)]
		if i != 0 {
			if ourMap[string(s[i-1])] < ourMap[string(v)] {
				val -= 2 * ourMap[string(s[i-1])]
			}
		}
	}
	return val
}