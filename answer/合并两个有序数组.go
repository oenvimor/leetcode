package answer

func Merge(nums1 []int, m int, nums2 []int, n int) {
	// 创建三个指针，分别指向 nums1 有效元素的末尾，nums2 有效元素的末尾， nums1 的末尾
	i := m - 1
	j := n - 1
	k := len(nums1) - 1
	// 如果 nums1 为空，直接把 nums2 所有元素放到 nums1 中
	if i < 0 {
		for j >= 0 {
			nums1[j] = nums2[j]
			j--
		}
		return
	}
	// 循环判断谁大谁小
	for i >= 0 && j >= 0 {
		// 谁大就把谁放到 nums1 的末尾
		if nums2[j] >= nums1[i] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
		k--
	}
	// 如果 nums2 中还有剩下的就放入 nums1 中
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}
