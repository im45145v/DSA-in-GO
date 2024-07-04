// Copied from the internet but learned
// Function to find two odd occurring numbers in an array
func twoOddNum(arr []int, n int) []int {
	// Initializing variables
	xor2 := arr[0] // to store the xor of all elements
	x := 0         // to store the first odd occurring number
	y := 0         // to store the second odd occurring number

	// Calculating xor of all elements in the array
	for i := 1; i < n; i++ {
		xor2 = xor2 ^ arr[i]
	}

	// Finding the rightmost set bit in xor2
	setBitNo := xor2 & ^(xor2 - 1)

	// Finding the first odd occurring number (x) and the second odd occurring number (y)
	for i := 0; i < n; i++ {
		if arr[i]&setBitNo != 0 {
			x = x ^ arr[i]
		} else {
			y = y ^ arr[i]
		}
	}

	// Returning the max and min values of x and y as a slice
	v := []int{max(x, y), min(x, y)}
	return v
}