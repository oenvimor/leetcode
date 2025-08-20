package answer

import "strings"

func ReverseWords(s string) string {
	if len(s) == 0 {
		return ""
	}
	res := make([]string, 0)
	start := 0
	idx := 0
	// 拿到第一个单词开始的下标 start
	for idx < len(s) {
		if s[idx] != ' ' {
			start = idx
			for idx < len(s) {
				if s[idx] == ' ' {
					res = append(res, s[start:idx])
					break
				}
				idx++
			}
			if idx >= len(s) {
				break
			}
		}
		idx++
	}
	if s[idx-1] != ' ' {
		res = append(res, s[start:])
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return strings.Replace(strings.Join(res, ","), ",", " ", -1)
}
