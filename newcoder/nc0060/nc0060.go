package nc0060

func generateParenthesis(n int) (ans []string) {
	var gen func(curr string, openLeft, closeLeft int)
	gen = func(curr string, openLeft, closeLeft int) {
		if openLeft == 0 && closeLeft == 0 {
			ans = append(ans, curr)
			return
		}
		if openLeft > closeLeft {
			return
		}
		if openLeft > 0 {
			gen(curr+"(", openLeft-1, closeLeft)
		}
		if closeLeft > 0 {
			gen(curr+")", openLeft, closeLeft-1)
		}
	}

	gen("", n, n)
	return
}
