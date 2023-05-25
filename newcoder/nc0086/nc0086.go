package nc0086

// 大数相加
func solve(s string, t string) (ans string) {
	i, j := len(s)-1, len(t)-1
	sum, carry := 0, 0
	var res []byte
	for i >= 0 || j >= 0 {
		sum = carry
		if i >= 0 {
			sum += int(s[i] - '0')
		}
		if j >= 0 {
			sum += int(t[j] - '0')
		}
		res = append(res, byte(sum%10+'0'))
		carry = sum / 10
		i--
		j--
	}

	if carry > 0 {
		res = append(res, '1')
	}

	for p, q := 0, len(res)-1; p < q; p, q = p+1, q-1 {
		res[p], res[q] = res[q], res[p]
	}

	ans = string(res)
	return
}
