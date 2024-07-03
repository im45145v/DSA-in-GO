func singleNumber(nums []int) int {
	x := nums[0]
	for _, i := range nums[1:] {
		x ^= i
	}
	return x
}