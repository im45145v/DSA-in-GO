package main

import "fmt"

func swapTwoNumbers(m, n int) []int {
	m = m ^ n
	n = n ^ m
	m = m ^ n
	return []int{m, n}
}
func main() {
	fmt.Println(swapTwoNumbers(2, 3))
}
