func generateParenthesis(n int) []string {
	var res []string
	backtrack(&res, n, 0, 0, "")
	return res
}
func backtrack(res *[]string, n, open, close int, curSeq string) {
	if open == close && close == n {
		*res = append(*res, curSeq)
		return
	}
	if open < n {
		curSeq += "("
		open++
		backtrack(res, n, open, close, curSeq)
		open--
		curSeq = curSeq[:len(curSeq)-1]
	}
	if close < open {
		curSeq += ")"
		close++
		backtrack(res, n, open, close, curSeq)
		close--
		curSeq = curSeq[:len(curSeq)-1]
	}
}

// From internet
// func generateParenthesis(n int) []string {
//     result := make([]string, 0)
//     generate(&result, "(", 1, 0, n)
//     return result
// }

// func generate(result *[]string, s string, open int, close int, max int) {
//     if (open + close) == (max * 2) {
//         *result = append(*result, s)
//         return
//     }
//     if open < max {
//         generate(result, s+"(", open+1, close, max)
//     }
//     if close < max && open > close {
//         generate(result, s+")", open, close+1, max)
//     }
// }