// Same as Koko banana problem
func smallestDivisor(nums []int, threshold int) int {
	l, r, mid, totalsum := 1, findMax(nums)+1, 0, 0
	for l <= r {
		mid = (l + r) / 2
		for _, k := range nums {
			totalsum += int(math.Ceil(float64(k) / float64(mid)))
		}
		if totalsum > threshold {
			l = mid + 1
		} else {
			r = mid - 1
		}
		totalsum = 0
	}
	return l

}
func findMax(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}