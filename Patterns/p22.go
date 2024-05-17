package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scanln(&i)
	var p, q int
	for row := 0; row < 2*i-1; row++ {
		for col := 0; col < 2*i-1; col++ {
			p = convert(row, i)
			q = convert(col, i)
			if p > q {
				fmt.Print(i - q)
			} else {
				fmt.Print(i - p)
			}
		}
		fmt.Println()
	}
}

func convert(val, ma int) int {
	if val >= ma {
		return 2*ma - val - 2
	}
	return val
}
