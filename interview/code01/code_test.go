package code01

import (
	"testing"
)

func TestCode02(t *testing.T) {
	ans := dfs([]string{"1", "2", "3", "4"})
	// [ "1", "2", "12" ]
	t.Log(ans)
}

func dfs(digits []string) (ans []string) {
	if len(digits) == 1 {
		ans = digits
		return
	}

	// 1 2 3 4

	for i := 0; i < len(digits); i++ {
		curr := digits[i]      // 1
		var newDigits []string // 2 3 4
		for j := 0; j < len(digits); j++ {
			if j == i {
				continue
			}
			newDigits = append(newDigits, digits[j])
		}
		// fmt.Println(newDigits)

		res := dfs(newDigits) //
		// fmt.Println(res)
		for _, r := range res {
			ans = append(ans, curr+r)
		}
	}

	return
}
