func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 1
	for k, v := range nums[1:] {
		if v != nums[k] {
			nums[i] = v
			i++
		}
	}
	return i
}
