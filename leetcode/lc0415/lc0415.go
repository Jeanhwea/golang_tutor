package lc0415

// 字符串相加
func addStrings(num1 string, num2 string) (ans string) {
	n, m := len(num1), len(num2)
	carry, sum, i, j := 0, 0, n-1, m-1

	for i >= 0 || j >= 0 {
		sum = carry
		if i >= 0 {
			sum += int(num1[i] - '0')
		}
		if j >= 0 {
			sum += int(num2[j] - '0')
		}
		ans = string(byte((sum%10)+'0')) + ans
		carry = sum / 10
		i--
		j--
	}

	if carry > 0 {
		ans = "1" + ans
	}

	return
}
