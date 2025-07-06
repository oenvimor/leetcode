package answer

func LongestPalindrome(s string) string {
	// 如果字符串长度小于 2 则直接返回
	if len(s) < 2 {
		return s
	}
	// 初始化起始位置
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		// 假设当前值为中心位置，并且判断是奇数和偶数两种情况
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		// 得到最长的字符串的长度
		maxLen := maxNum(len1, len2)
		if maxLen > end-start+1 {
			// 如果长度超过上一次记录的长度则需要更新起始位置
			start = i - (maxLen-1)/2
			end = maxLen/2 + i
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left int, right int) int {
	// 如果左右指针没有超出界限，并且左右指针所指位置的值相等，则向左右扩展，直到不满足条件就返回长度
	var length int
	for left >= 0 && right < len(s) && s[left] == s[right] {
		length = right - left + 1
		left--
		right++
	}
	return length
}

func maxNum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
