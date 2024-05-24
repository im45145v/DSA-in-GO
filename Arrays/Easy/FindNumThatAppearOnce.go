func singleNumber(nums []int) int {
	for _, n := range nums[1:] {
		nums[0] ^= n
	}
	return nums[0]
}