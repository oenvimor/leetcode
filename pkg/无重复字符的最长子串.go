package pkg

func lengthOfLongestSubstring(s string) int {
	charIdxMap := make(map[byte]int)
	start := 0
	maxLen := 0
	for i := 0; i < len(s); i++ {
		if lastIdx, found := charIdxMap[s[i]]; found && start <= lastIdx {
			start = lastIdx + 1
		}
		charIdxMap[s[i]] = i
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
	}
	return maxLen
}
