//Not coded by me and took from leetcode which is way better
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	var found = -1
	for left <= right {
		mid := (left + right) / 2

		if nums[mid] > target {
			right = mid - 1
		}

		if nums[mid] < target {
			left = mid + 1
		}

		if nums[mid] == target {
			found = mid
			break
		}
	}

	if found == -1 {
		return []int{-1, -1}
	}

	first, last := found, found

	for i := found - 1; i >= 0; i-- {
		if nums[i] != target {
			break
		}
		if nums[i] == target {
			first = i
		}
	}

	for i := found + 1; i < len(nums); i++ {
		if nums[i] != target {
			break
		}
		if nums[i] == target {
			last = i
		}
	}
	return []int{first, last}
}