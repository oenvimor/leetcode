package answer

import (
	"strconv"
	"strings"
)

func Multiply(num1 string, num2 string) string {
	res := make([]int, len(num1)+len(num2))
	tmp := make([]string, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			m := int(num1[i] - '0')
			n := int(num2[j] - '0')
			sum := res[i+j+1] + m*n
			res[i+j+1] = sum % 10
			res[i+j] += sum / 10
		}
	}
	var i int
	for i = 0; i < len(res)-1; i++ {
		if res[i] != 0 {
			break
		}
	}
	res = res[i:]
	for idx, v := range res {
		tmp[idx] = strconv.Itoa(v)
	}
	return strings.Join(tmp, "")
}
