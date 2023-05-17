package nc0058

import "sort"

func Permutation(str string) (ans []string) {
	chars := []byte(str)
	visited := make([]bool, len(chars), len(chars))
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	dfs(&ans, chars, []byte{}, visited)
	return
}

func dfs(res *[]string, chars, choose []byte, visited []bool) {
	if len(choose) == len(chars) {
		*res = append(*res, string(choose))
		return
	}

	for i := 0; i < len(chars); i++ {
		// 如果 chars[i] 已经添加过了则不需要添加
		if visited[i] {
			continue
		}

		// 注意 chars 是排好序的, 这里对应的测试用例
		//   a1 a2 a3 b
		//
		//   a1
		//   |- a1 a2
		//   |- a1 a3 (这种是重复情形, 需要去重)
		//   |- a1 b
		//
		if i > 0 && chars[i-1] == chars[i] && !visited[i-1] {
			continue
		}

		visited[i], choose = true, append(choose, chars[i])
		dfs(res, chars, choose, visited)
		visited[i], choose = false, choose[:len(choose)-1]
	}
}
