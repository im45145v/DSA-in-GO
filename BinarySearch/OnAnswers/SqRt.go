package main

import (
	"fmt"
)

func sqrt(num int) int {
	if num == 1 {
		return 1
	}
	l, r := 0, num
	var mid int
	for l <= r {
		mid = (l + r) / 2
		if mid*mid > num {
			r = mid - 1
		} else {
			l = mid + 1
		}

	}
	return r
	// Bottom one can also be used
	// for l <= r {
	// 	mid = (l + r) / 2
	// 	if mid*mid >= num {
	// 		r = mid - 1
	// 	} else {
	// 		l = mid + 1
	// 	}
	// }
	// if mid*mid == num {
	// 	return mid
	// } else {
	// 	return mid - 1
	// }
}

func main() {
	//Your custom input
	n := 49
	fmt.Print(sqrt(n))
}
