package main

import "fmt"

func isSet(n int, k int) bool {
	mask := 1 << (k - 1)
	// fmt.Println(n)
	// fmt.Println(mask)
	// fmt.Println(n & mask)
	return (n & mask) != 0
}
func main() {
	if isSet(75, 4) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
