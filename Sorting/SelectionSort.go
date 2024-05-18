package main

import "fmt"

func main() {
	nums := []int{5, 2, 9, 3, 6, 7, 1, 4, 8, 3}
	l := len(nums)
	var i, j, k int
	for i = 0; i < l-1; i++ {
		k = i
		for j = i; j < l; j++ {
			if nums[j] < nums[k] {
				k = j
			}
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
	fmt.Print(nums)
}
