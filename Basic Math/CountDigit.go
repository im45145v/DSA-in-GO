package main

import (
	"fmt"
)

func main() {
	var inp, count int
	fmt.Scanln(&inp)
	for count = 0; inp > 0; inp /= 10 {
		count++
	}
	fmt.Println(count)
}

// according to video (the PS changed in coding platform they mentioned)
