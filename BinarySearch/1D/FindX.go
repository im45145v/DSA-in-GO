func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start, end := 0, len(nums)-1
	var mid int
	for start <= end {
		mid = (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}