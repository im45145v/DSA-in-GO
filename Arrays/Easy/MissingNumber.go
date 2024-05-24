func missingNumber(nums []int) int {
	sumo := (len(nums) * (len(nums) + 1)) / 2
	for _, n := range nums {
		sumo -= n
	}
	return sumo
}