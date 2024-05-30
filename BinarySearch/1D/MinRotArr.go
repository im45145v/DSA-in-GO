func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	least := nums[0]
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] < least {
			least = nums[mid]
		}
		if nums[low] <= nums[mid] {
			if nums[low] < least {
				least = nums[low]
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return least
}
