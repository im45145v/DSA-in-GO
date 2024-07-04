package main

import "fmt"

func xorl2r(i, j int) int {
	return f1(i-1) ^ f1(j)
}
func f1(j int) int {
	if j%4 == 1 {
		return 1
	} else if j%4 == 2 {
		return j + 1
	} else if j%4 == 3 {
		return 0
	} else {
		return j
	}
}
func main() {
	fmt.Print(xorl2r(4, 8))
}
