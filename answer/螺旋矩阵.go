package answer

func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	ans := make([]int, 0)
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	for top <= bottom && left <= right {
		// 从左往右
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
		}
		top++
		// 从上往下
		for i := top; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--
		// 从右往左
		// 判断是否有剩余行
		if top <= bottom {
			for i := right; i >= left; i-- {
				ans = append(ans, matrix[bottom][i])
			}
			bottom--
		}
		// 从下往上
		// 判断是否有剩余列
		if left <= right {
			for i := bottom; i >= top; i-- {
				ans = append(ans, matrix[i][left])
			}
			left++
		}
	}
	return ans
}
