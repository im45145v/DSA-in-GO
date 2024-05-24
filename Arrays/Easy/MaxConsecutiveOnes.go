func findMaxConsecutiveOnes(nums []int) int {
	var cm, am int
	for _, n := range nums {
		if n == 1 {
			cm++
		} else {
			if cm > am {
				am = cm
			}
			cm = 0
		}
	}
	if cm > am {
		am = cm
	}
	return am
}