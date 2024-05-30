func singleNonDuplicate(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := (low + high) / 2
		if mid%2 == 1 {
			mid--
		}
		if nums[mid] != nums[mid+1] {
			high = mid
		} else {
			low = mid + 2
		}
	}
	return nums[low]
}

// There are more solutions here https://leetcode.com/problems/single-element-in-a-sorted-array/solutions/3869802/go-multiple-solutions-in-go-golang-100