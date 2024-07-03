func minBitFlips(start int, goal int) int {
	x, count := start^goal, 0
	for ; x > 0; count++ {
		x = x & (x - 1)
	}
	return count
}