package main

import "fmt"

func rotate(nums []int, k int) {
	l := len(nums)
	if l == 0 {
		return
	}
	k = k % l
	if k == 0 {
		return
	}
	reverse(nums, 0, k-1)
	reverse(nums, k, l-1)
	reverse(nums, 0, l-1)
}
func reverse(nums []int, s int, e int) {
	for s < e {
		nums[s], nums[e] = nums[e], nums[s]
		s++
		e--
	}
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Print(nums)
}
