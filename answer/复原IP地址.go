package answer

import (
	"strings"
)

func RestoreIpAddresses(s string) []string {
	var res []string
	var path []string

	var backtrack func(int)
	backtrack = func(start int) {
		if len(path) == 4 {
			if start == len(s) {
				res = append(res, strings.Join(path, "."))
			}
			return
		}
		for i := 1; i <= 3 && start+i <= len(s); i++ {
			sample := s[start : start+i]
			if isValid(sample) {
				path = append(path, sample)
				backtrack(start + i)
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(0)
	return res
}

func isValid(s string) bool {
	val := 0
	if len(s) > 3 {
		return false
	}
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	for i := range s {
		val = val*10 + int(s[i]-'0')
	}
	return val >= 0 && val <= 255
}
