package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	fmt.Scanln(&a, &b)
	for a != 0 && b != 0 {
		if a > b {
			a = a % b
		} else {
			b = b % a
		}
	}
	fmt.Print(math.Max(float64(a), float64(b)))
}
