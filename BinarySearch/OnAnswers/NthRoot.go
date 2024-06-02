package main

import (
	"fmt"
	"math"
)

func Nthroot(num int, n int) int {
	if num == 1 {
		return 1
	}
	l, r := 0, num
	var mid int
	for l <= r {
		mid = (l + r) / 2
		if int(math.Pow(float64(mid), float64(n))) >= num {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if int(math.Pow(float64(mid), float64(n))) == num {
		return mid
	} else {
		return -1
	}
}

func main() {
	//Your custom input
	num, n := 80, 4
	fmt.Print(Nthroot(num, n))
}
