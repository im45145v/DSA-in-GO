package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	var p, q int
	q = x
	for x > 0 {
		p = p*10 + x%10
		x /= 10
	}
	return p == q
}
func main() {
	var x int
	fmt.Scan(&x)
	fmt.Print(isPalindrome(x))
}
