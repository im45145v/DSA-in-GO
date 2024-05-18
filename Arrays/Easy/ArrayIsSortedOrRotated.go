func check(nums []int) bool {
	var track int
	for i, v := range nums {
		if v > nums[(i+1)%len(nums)] {
			track++
		}
		if track > 1 {
			return false
		}
	}
	return true
}