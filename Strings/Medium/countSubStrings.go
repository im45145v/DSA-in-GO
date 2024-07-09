package main

import (
	"fmt"
)

func substrCount(s string, k int) int64 {
	n := len(s)
	var res int64 = 0

	cnt := make([]int, 26)

	for i := 0; i < n; i++ {
		distCount := 0
		for j := i; j < n; j++ {
			if cnt[int(s[j]-'a')] == 0 {
				distCount++
			}
			cnt[int(s[j]-'a')]++
			if distCount == k {
				res++
			}
			if distCount > k {
				break
			}
		}
		for j := 0; j < 26; j++ {
			cnt[j] = 0
		}
	}

	return res
}

func main() {
	s := "abaaca"
	k := 1
	fmt.Println(substrCount(s, k)) // Output: 10
}
