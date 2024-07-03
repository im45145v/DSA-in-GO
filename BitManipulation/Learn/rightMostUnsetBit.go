package main

import "fmt"

func rightMostUnsetBitFlip(n int) int {
	if n&(n+1) == 0 {
		return n
	}
	x := (n | (n + 1))
	return x
}
func main() {
	fmt.Println(rightMostUnsetBitFlip(2))
}
