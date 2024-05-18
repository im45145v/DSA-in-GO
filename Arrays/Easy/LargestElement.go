package main

import "fmt"

func main() {
	nums := []int{5, 2, 9, 3, 6, 7, 1, 4, 8, 3}
	fmt.Print(maxele(nums))
}
func maxele(nums []int) int {
	var x int
	for _, v := range nums {
		if v > x {
			x = v
		}
	}
	return x
}
