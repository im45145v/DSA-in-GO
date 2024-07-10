func beautySum(s string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		count := make([]int, 26)
		for j := i; j < len(s); j++ {
			count[s[j]-'a']++
			ans += max(count) - getMinFreq(count)
		}
	}
	return ans
}

func max(count []int) int {
	maxFre := 0
	for _, fre := range count {
		if fre > maxFre {
			maxFre = fre
		}
	}
	return maxFre
}
func getMinFreq(count []int) int {
	minFreq := math.MaxInt32
	for _, fre := range count {
		if fre < minFreq && fre > 0 {
			minFreq = fre
		}
	}
	return minFreq
}