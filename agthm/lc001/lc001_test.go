// Copyright 2023 The Go Authors in ByteDance Inc. All rights reserved.
// Use of this source code should follow terms of ByteDance Inc.
//
// Authors:     hujinghui.jeffrey@bytedance.com
// Create Date: 2023-04-29
// Update Date: 2023-04-29
//

package lc001

import "testing"

func Test_LC001_01(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	ans := twoSum(nums, target)
	t.Log(ans)
}
