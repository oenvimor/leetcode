package answer

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	nums1 := strings.Split(version1, ".")
	nums2 := strings.Split(version2, ".")
	maxLen := Max(len(nums1), len(nums2))
	for i := 0; i < maxLen; i++ {
		var s1, s2 int
		if i < len(nums1) {
			s1, _ = strconv.Atoi(nums1[i])
		} else {
			s1 = 0
		}
		if i < len(nums2) {
			s2, _ = strconv.Atoi(nums2[i])
		} else {
			s2 = 0
		}
		if s1 < s2 {
			return -1
		} else if s1 > s2 {
			return 1
		}
	}
	return 0
}
