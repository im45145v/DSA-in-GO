package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scan(&i)
	p := 1

	for row := 0; row < i+1; row++ {
		for col := 0; col < row; col++ {
			if p == 0 {
				p = 1
			} else {
				p = 0
			}
			fmt.Print(p)
		}
		fmt.Println()
		if row%2 == 0 {
			if p == 0 {
				p = 1
			} else {
				p = 0
			}

		}
	}
}
