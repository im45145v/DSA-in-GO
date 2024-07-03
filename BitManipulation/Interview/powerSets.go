func subsets(nums []int) [][]int {
	n := len(nums)
	var result [][]int
	for i := 0; i < (1 << n); i++ {
		subSet := []int{}
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				subSet = append(subSet, nums[j])
			}
		}
		result = append(result, subSet)
	}
	return result
}