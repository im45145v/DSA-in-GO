package main

import (
	"fmt"
)

func main() {
	var i, j int
	fmt.Scan(&i)
	for i > 0 {
		j *= 10
		j += i % 10
		i /= 10
	}
	fmt.Println(j)

}

// This code is for leetcode as there is limits 
// func reverse(x int) int {
// 	var p, j int
// 	if x < 0 {
// 		x *= -1
// 		p = 1
// 	}
// 	for x > 0 {
// 		pop := x % 10
// 		x /= 10
// 		if j > math.MaxInt32/10 || (j == math.MaxInt32/10 && pop > 7) {
// 			return 0
// 		}
// 		if j < math.MinInt32/10 || (j == math.MinInt32/10 && pop < -8) {
// 			return 0
// 		}
// 		j = int(int32(j)*10 + int32(pop))
// 	}
// 	if p == 1 {
// 		return -j
// 	}
// 	return j
// }