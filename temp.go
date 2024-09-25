package main

import (
	"fmt"
)

func findComp(s string) int {
	n := len(s)
	comp := make([]int, n)
	for i := range comp {
		comp[i] = -1
	}
	count := 0

	for i := 0; i < n; i++ {
		if s[i] == 's' {
			if i > 0 && s[i-1] == '.' {
				if comp[i-1] == -1 {
					comp[i-1] = i
					count++
					if i-3 >= 0 && s[i-2] == 's' && comp[i-3] == i-2 {
						count--
					}
				}
			} else if i < n-1 && s[i+1] == '.' {
				if comp[i+1] == -1 {
					comp[i+1] = i
					count++
				}
			} else {
				return -1
			}
		}
	}
	return count
}

func main() {
	testCases := []string{
		"ss.s.s.s.s",
		"s.s.s.s.ss",
		"s.ss.ss.ss.s.s.s",
		"s.s.sss.s.s.s...",
		"s.ss.ss.s....s.ss.ss.",
	}

	for _, testCase := range testCases {
		result := findComp(testCase)
		fmt.Printf("Input: %s => Output: %d\n", testCase, result)
	}
}
