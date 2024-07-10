package main

import "fmt"

func nextPermutation(nums []int) {
	p, j, k := len(nums)-1, len(nums)-1, len(nums)-1

	for p > 0 && nums[p] <= nums[p-1] {
		p--
	}

	if p == 0 {
		reverse(nums)
		return
	}

	for nums[j] <= nums[p-1] {
		j--
	}

	nums[p-1], nums[j] = nums[j], nums[p-1]

	for p < k {
		nums[p], nums[k] = nums[k], nums[p]
		p++
		k--
	}
}

func reverse(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println("Original nums:", nums)

	nextPermutation(nums)
	fmt.Println("Next permutation:", nums)
}
