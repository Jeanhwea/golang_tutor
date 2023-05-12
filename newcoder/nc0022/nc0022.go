package nc0022

// 比较版本号
func compare(version1 string, version2 string) (ans int) {
	n, m := len(version1), len(version2)
	for i, j := 0, 0; i < n || j < m; {
		num1 := 0
		for ; i < n && version1[i] != '.'; i++ {
			num1 = 10*num1 + int(byte(version1[i])-'0')
		}
		num2 := 0
		for ; j < m && version2[j] != '.'; j++ {
			num2 = 10*num2 + int(byte(version2[j]-'0'))
		}

		if num1 > num2 {
			ans = 1
			return
		} else if num1 < num2 {
			ans = -1
			return
		} else {
			i++
			j++
		}
	}
	return
}
