package answer

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	// 先判断奇偶
	if (m+n)%2 == 1 {
		return float64(getKth(nums1, nums2, (m+n)/2+1))
	} else {
		first := getKth(nums1, nums2, (m+n)/2)
		second := getKth(nums1, nums2, (m+n)/2+1)
		return float64(first+second) / 2
	}
}

func getKth(nums1, nums2 []int, k int) int {
	idx1, idx2 := 0, 0
	for {
		if idx1 == len(nums1) {
			return nums2[idx2+k-1]
		}
		if idx2 == len(nums2) {
			return nums1[idx1+k-1]
		}
		if k == 1 {
			return Min(nums1[idx1], nums2[idx2])
		}
		half := k / 2
		newIdx1 := Min(idx1+half, len(nums1)) - 1
		newIdx2 := Min(idx2+half, len(nums2)) - 1
		if nums1[newIdx1] <= nums2[newIdx2] {
			k -= newIdx1 - idx1 + 1
			idx1 = newIdx1 + 1
		} else {
			k -= newIdx2 - idx2 + 1
			idx2 = newIdx2 + 1
		}
	}
}

//func Min(a, b int) int {
//	if a > b {
//		return b
//	} else {
//		return a
//	}
//}
