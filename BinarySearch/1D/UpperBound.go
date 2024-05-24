package main

import (
	"fmt"
	"sort"
)

// UpperBound finds the first position in a sorted array where the value is greater than the target
func UpperBound(nums []int, tar int) int {
	if nums[0] > tar {
		return -1
	}
	l, r, ans := 0, len(nums)-1, len(nums)
	var mid int
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] > tar {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}

func main() {
	nums := []int{5, 6, 8, 9, 6, 5, 5, 6}
	low := 7
	sort.Ints(nums) // Sorting the array
	fmt.Print(UpperBound(nums, low))
}
