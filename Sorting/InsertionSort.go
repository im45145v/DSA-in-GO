package main

import "fmt"

func main() {
	nums := []int{5, 2, 9, 3, 6, 3, 7, 1, 4, 8}
	l := len(nums)
	var i, j int
	for i = 0; i < l; i++ {
		j = i
		for j > 0 && nums[j-1] > nums[j] {
			nums[j-1], nums[j] = nums[j], nums[j-1]
			j--
		}

	}
	fmt.Print(nums)
}
