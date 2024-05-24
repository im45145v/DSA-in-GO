package main

import "fmt"

func LowBou(nums []int, leng int, tar int) int {
	if nums[0] > tar {
		return -1
	}
	l, r := 0, leng-1
	var mid, ans int
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] >= tar {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}

	}
	return ans
}

func main() {
	nums := []int{1, 2, 8, 10, 11, 12, 19}
	low := 5
	fmt.Print(LowBou(nums, len(nums), low))
}
