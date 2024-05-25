package main

import "fmt"

func Occur(nums []int, target int) int {
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
		return 0
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
	return last - first + 1
}
func main() {
	nums := []int{5, 7, 7, 8, 8, 8, 10}
	target := 9
	fmt.Println(Occur(nums, target))
}
