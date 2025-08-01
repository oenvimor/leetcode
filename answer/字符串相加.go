package answer

func AddStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	carry := 0
	var results []byte
	for i >= 0 || j >= 0 || carry > 0 {
		var n1, n2 int
		if i >= 0 {
			n1 = int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			n2 = int(num2[j] - '0')
			j--
		}
		sum := n1 + n2 + carry
		carry = sum / 10
		results = append([]byte{byte(sum%10 + '0')}, results...)
	}
	return string(results)
}
