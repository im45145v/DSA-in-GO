func searchInsert(nums []int, target int) int {
	start, end, mid := 0, len(nums)-1, 0
	for start <= end {
		mid = (start + end) / 2
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid + -1
		}
	}
	return start
}