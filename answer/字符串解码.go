package answer

import "unicode"

func DecodeString(s string) string {
	stackNum := make([]int, 0)
	stackStr := make([]string, 0)
	currNum := 0
	currStr := ""
	for _, ch := range s {
		if unicode.IsDigit(ch) {
			currNum = currNum*10 + int(ch-'0')
		} else if ch == '[' {
			stackNum = append(stackNum, currNum)
			stackStr = append(stackStr, currStr)
			currNum = 0
			currStr = ""
		} else if ch == ']' {
			n := stackNum[len(stackNum)-1]
			stackNum = stackNum[:len(stackNum)-1]
			prevStr := stackStr[len(stackStr)-1]
			stackStr = stackStr[:len(stackStr)-1]
			tmp := ""
			for i := 0; i < n; i++ {
				tmp += currStr
			}
			currStr = prevStr + tmp
		} else {
			currStr += string(ch)
		}
	}
	return currStr
}
