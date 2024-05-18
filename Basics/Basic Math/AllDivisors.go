package main

import (
	"fmt"
	"math"
	"sort"
)

func printDivisors(n int) {
	var divisors []int
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if i != n/i {
				divisors = append(divisors, n/i)
			}
		}
	}
	sort.Ints(divisors)
	fmt.Print(divisors)
}
func main() {
	var num int
	fmt.Scanln(&num)
	printDivisors(num)
}
