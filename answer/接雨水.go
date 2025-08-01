package answer

//func Trap(height []int) int {
//	sum := 0
//	for i := 1; i < len(height)-1; i++ {
//		maxLeft, maxRight := 0, 0
//		// 往左找最大值
//		for j := 0; j < i; j++ {
//			if height[j] > maxLeft {
//				maxLeft = height[j]
//			}
//		}
//		// 往右找最大值
//		for k := i + 1; k < len(height); k++ {
//			if height[k] > maxRight {
//				maxRight = height[k]
//			}
//		}
//		// 找最大值中较小的数
//		minH := Min(maxLeft, maxRight)
//		// 遍历每个格子并计算当前格子能存储多少税
//		if minH > height[i] {
//			sum += minH - height[i]
//		}
//	}
//	return sum
//}

func Trap(height []int) int {
	sum := 0
	n := len(height)

	leftMax := make([]int, n)
	rightMax := make([]int, n)

	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = Max(leftMax[i-1], height[i])
	}

	rightMax[n-1] = height[n-1]
	for j := n - 2; j >= 0; j-- {
		rightMax[j] = Max(rightMax[j+1], height[j])
	}

	for k := 1; k < n-1; k++ {
		sum += Min(leftMax[k], rightMax[k]) - height[k]
	}
	return sum
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

//func Max(a, b int) int {
//	if a > b {
//		return a
//	} else {
//		return b
//	}
//}
