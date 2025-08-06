package answer

func GenerateParenthesis(n int) []string {
	var res []string
	var backtrack func(path string, left, right int)
	backtrack = func(path string, left, right int) {
		if left == n && right == n {
			res = append(res, path)
			return
		}
		if left < n {
			backtrack(path+"(", left+1, right)
		}
		if right < left {
			backtrack(path+")", left, right+1)
		}
	}
	backtrack("", 0, 0)
	return res
}
