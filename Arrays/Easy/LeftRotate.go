package main

import "fmt"

func Rotate(nums []int, k int) {
	temp := nums[0]
	for i := 0; i < k-1; i++ {
		nums[i] = nums[i+1]
	}
	nums[k-1] = temp
	fmt.Print(nums)
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	Rotate(nums, len(nums))
}
