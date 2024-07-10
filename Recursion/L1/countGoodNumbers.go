const mod int64 = 1e9 + 7

func countGoodNumbers(n int64) int {
	memo := make(map[int64]int64)
	return int(helper(n, memo))
}

func helper(n int64, memo map[int64]int64) int64 {
	if n == 0 {
		return 1
	}
	if val, found := memo[n]; found {
		return val
	}

	var result int64
	if n%2 == 0 {
		result = (4 * helper(n-1, memo)) % mod
	} else {
		result = (5 * helper(n-1, memo)) % mod
	}

	memo[n] = result
	return result
}