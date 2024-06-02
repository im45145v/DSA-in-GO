func minEatingSpeed(piles []int, h int) int {
	max := findMax(piles)
	l, r, mid, total := 1, max, 0, 0
	for l <= r {
		mid = (l + r) / 2
		total = 0
		for _, pile := range piles {
			total += int(math.Ceil(float64(pile) / float64(mid)))
		}
		if total > h {
			l = mid + 1
		} else {
			r = mid - 1
		}
		total = 0
	}
	return l
}
func findMax(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}