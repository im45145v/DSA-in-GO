package main

import "fmt"

func main() {
	nums := []int{5, 2, 9, 3, 6, 7, 1, 4, 8, 3}
	fmt.Print(maxele(nums))
}
func maxele(nums []int) int {
	var Largest, SecongLargest int
	for _, value := range nums {
		if nums[value] > Largest {
			Largest, SecongLargest = nums[value], Largest

		} else if SecongLargest < nums[value] && nums[value] < Largest {
			SecongLargest = nums[value]
		}
	}
	return SecongLargest
}
