package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	lb := LowBou(nums, len(nums), target)
	if lb == len(nums) || nums[lb] != target {
		return []int{-1, -1}
	}
	return []int{lb, UpperBound(nums, target) - 1}
}

func LowBou(nums []int, leng int, tar int) int {
	l, r := 0, leng-1
	var mid int
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] >= tar {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func UpperBound(nums []int, tar int) int {
	l, r := 0, len(nums)-1
	var mid int
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] > tar {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange(nums, target))
}
