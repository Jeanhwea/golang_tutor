package nc0077

// 最长的括号子串 - 单调栈解法
func longestValidParentheses2(s string) (ans int) {
	stack := []int{-1}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else { // s[i] == ')'
			if len(stack) == 1 { // 发现括号不匹配
				stack[0] = i
				continue
			}

			stack = stack[:len(stack)-1]
			length := i - stack[len(stack)-1]
			if length > ans {
				ans = length
			}
		}
	}
	return
}

// 最长的括号子串 - 动规解法
func longestValidParentheses(s string) (ans int) {
	dp := make([]int, len(s))
	dp[0] = 0
	for i := 1; i < len(s); i++ {
		if s[i] == '(' {
			dp[i] = 0
		} else { // s[i] == ')'
			if i-1 < 0 {
				continue
			}
			if s[i-1] == '(' {
				if i-2 >= 0 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else { // s[i-1] == ')'
				j := i - dp[i-1] - 1 // j 是跳过 s[i-1] 匹配的括号字串后的下标
				if j >= 0 && s[j] == '(' {
					if j-1 >= 0 {
						dp[i] = dp[j-1] + dp[i-1] + 2
					} else {
						dp[i] = dp[i-1] + 2
					}
				}
			}
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return
}
