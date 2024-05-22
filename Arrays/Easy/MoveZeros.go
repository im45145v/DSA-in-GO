func moveZeroes(nums []int) {
	l := len(nums)
	if l == 0 {
		return
	}
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
	}
	for j := i; j < l; j++ {
		nums[j] = 0
	}

}