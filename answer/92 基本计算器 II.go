package answer

import "unicode"

func Calculate(s string) int {
	sign := byte('+')
	stack := make([]int, 0)
	num := 0
	for i := 0; i < len(s); i++ {
		ch := s[i]
		// 累计数字
		if unicode.IsDigit(rune(ch)) {
			num = num*10 + int(ch-'0')
		}
		// 遇到符号或到末尾结算
		if (!unicode.IsDigit(rune(ch)) && ch != ' ') || i == len(s)-1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, top*num)
			case '/':
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, top/num)
			}
			sign = ch
			num = 0
		}

	}
	res := 0
	for _, v := range stack {
		res += v
	}
	return res
}
