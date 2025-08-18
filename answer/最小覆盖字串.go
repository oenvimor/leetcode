package answer

const (
	s = "ADOBECODEBANC"
	t = "ABC"
)

func MinWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	if len(t) > len(s) {
		return ""
	}
	// 创建 map 用以记录字符出现频次
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	// 创建 map 用于记录 s 中字符出现频次
	window := make(map[byte]int)
	// 左右指针用来控制滑动窗口大小
	left, right := 0, 0
	// 有效指针用来判断是否已经覆盖
	valid := 0
	start, length := 0, len(s)+1
	// 右指针先开始移动
	for right < len(s) {
		// 取到右指针所在位置的字符
		c := s[right]
		// 右指针移动
		right++
		// 窗口 map 记录字符出现次数
		window[c]++
		// 如果窗口中出现该字符的次数等于 need 中的字符出现的次数 valid 增加数量
		if window[c] == need[c] {
			valid++
		}
		// 如果满足覆盖条件则左指针开始控制窗口收缩
		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}
			d := s[left]
			left++
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	if length > len(s) {
		return ""
	}
	return s[start : start+length]
}
