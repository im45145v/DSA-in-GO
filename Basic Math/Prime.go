package main

import (
	"fmt"
	"math"
)

func isprime(n int) bool {
	if n == 2 {
		return true
	}
	if n <= 1 || n%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	var num int
	fmt.Scanln(&num)
	fmt.Print(isprime(num))
}
