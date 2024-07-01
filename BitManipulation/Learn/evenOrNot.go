package main

import "fmt"

func evenOrNot(n int) bool {
	return (n & 1) != 0
}
func main() {
	if evenOrNot(80) {
		fmt.Println("False")
	} else {
		fmt.Println("True")
	}
}
