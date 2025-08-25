package answer

func LongestCommonPrefix(strs []string) string {
	res := ""
	point := strs[0]
	for i := range point {
		for n := 1; n < len(strs); n++ {
			if i > len(strs[n])-1 || strs[n][i] != point[i] {
				return res
			}
		}
		res += string(point[i])
	}
	return res
}
