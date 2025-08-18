package answer

func LongestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}
	stack := []int{-1}
	maxLen := 0
	for i, v := range s {
		// 如果是 ( 则将下标入栈
		if v == '(' {
			stack = append(stack, i)
		} else {
			// 否则出栈，消费一个 (
			stack = stack[:len(stack)-1]
			// 判空，为空则将当前下标作为基准
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				// 若不为空，则计算长度
				length := i - stack[len(stack)-1]
				if length > maxLen {
					maxLen = length
				}
			}
		}
	}
	return maxLen
}
