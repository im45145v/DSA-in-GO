func merge(nums1 []int, m int, nums2 []int, n int) {
	// Corrected for loop
	for i, j := m-1, 0; i >= 0 && j < n; i, j = i-1, j+1 {
		if nums1[i] > nums2[j] {
			nums1[i], nums2[j] = nums2[j], nums1[i]
		} else {
			break
		}
	}
	quickSort(nums1, 0, m-1)
	quickSort(nums2, 0, n-1)
	copy(nums1[m:], nums2[:])
}
func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := (low - 1)

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return (i + 1)
}
func quickSort(arr []int, low int, high int) {
	if low < high {
		pi := partition(arr, low, high)

		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}