package main

import "fmt"

func Lengthk(nums []int, tar int) int {
	hm := make(map[int]int)
	var sum, maxl int
	for i, j := range nums {
		sum += j
		if sum == tar {
			maxl = i + 1
		}
		if _, there := hm[sum]; !there {
			hm[sum] = i
		}
		if _, there := hm[sum-tar]; there {
			if maxl < (i - hm[sum-tar]) {
				maxl = i - hm[sum-tar]
			}
		}
	}
	return maxl
}

func main() {
	nums := []int{10, 5, 2, 7, 1, 9}
	fmt.Print(Lengthk(nums, 15))
}
