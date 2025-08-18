package answer

func MySqrt(x int) int {
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)>>1
		if mid*mid > x {
			right = mid - 1
		} else if mid*mid < x {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
