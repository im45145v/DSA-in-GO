package main

import (
	"fmt"
	"strings"
)

func main() {
	var i int
	fmt.Scan(&i)
	x := strings.Repeat("*", i) + "\n"
	x = strings.Repeat(x, i)
	fmt.Println(x)
}
