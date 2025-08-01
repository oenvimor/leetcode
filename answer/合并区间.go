package answer

import "sort"

func MergeList(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		current := intervals[i]
		if current[0] > last[1] {
			merged = append(merged, current)
		} else {
			if current[1] > last[1] {
				last[1] = current[1]
			}
		}
	}
	return merged
}
