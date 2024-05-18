package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Print(fact(x))
}
func fact( p int) int {
	if p == 1 || p == 0 {
		return 1
	}
	return p * fact(p-1)
}
