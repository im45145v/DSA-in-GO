package main

import (
	"fmt"
	"math"
)

func isArmstrong(x int) bool {
	var p, q int
	q = x
	for x > 0 {
		p += int(math.Pow(float64(x%10), 3))
		x /= 10
	}
	return p == q
}
func main() {
	var x int
	fmt.Scan(&x)
	fmt.Print(isArmstrong(x))
}
