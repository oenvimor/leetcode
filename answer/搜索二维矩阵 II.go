package answer

func SearchMatrix(matrix [][]int, target int) bool {
	m, n := 0, len(matrix[0])-1
	for m < len(matrix) && n >= 0 {
		point := matrix[m][n]
		if target < point {
			n--
		} else if target > point {
			m++
		} else {
			return true
		}
	}
	return false
}
