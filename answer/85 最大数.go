package answer

import (
	"fmt"
	"sort"
	"strings"
)

func LargestNumber(nums []int) string {
	res := make([]string, len(nums))
	for i, n := range nums {
		res[i] = fmt.Sprintf("%d", n)
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i]+res[j] > res[j]+res[i]
	})
	r := strings.Join(res, "")
	if r[0] == '0' {
		return "0"
	}
	return r
}
