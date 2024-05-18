package main

import "fmt"

func main() {
	nums := []int{5, 2, 9, 3, 6, 7, 1, 4, 8, 3}
	l := len(nums)
	var i, j int
	for i = l; i >= 0; i-- {
		for j = 0; j < i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	fmt.Print(nums)
}
