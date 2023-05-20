package nc0066

import (
	"fmt"
)

// 最长公共子串
func LCS(str1 string, str2 string) (ans string) {
	n, m := len(str1), len(str2)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if str1[i] != str2[j] {
				continue
			}
			len := 0
			for ; i+len < n && j+len < m && str1[i+len] == str2[j+len]; len++ {
			}
			fmt.Printf("len=%v,  %v, %v\n", len, str1[i:i+len], str2[j:j+len])
		}
	}
	return
}

// 暴力解法
func LCS2(str1 string, str2 string) (ans string) {
	n, m := len(str1), len(str2)
	// 取 str1[i] 固定
	for i := 0; i < n; i++ {
		// 遍历 str2, 找到与 str[i] 相同的字符作为起点进行搜索
		for j := 0; j < m; j++ {
			if str1[i] != str2[j] {
				continue
			}

			// 搜索得到 str1 和 str2 的相同长度的字符串
			length := 0
			for ; i+length < n && j+length < m && str1[i+length] == str2[j+length]; length++ {
			}
			if length > len(ans) {
				ans = str1[i : i+length]
			}
		}
	}
	return
}
