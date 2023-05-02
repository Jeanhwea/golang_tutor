package lc088

// 合并两个有序数组
//
//    1. 这里需要使用一个临时数组 nums3 来保存 nums1 的变量
//
func merge2(nums1 []int, m int, nums2 []int, n int) {
	nums3 := make([]int, m)
	copy(nums3, nums1[:m])

	i, j, k := 0, 0, 0
	for ; i < m && j < n; k++ {
		if nums3[i] <= nums2[j] {
			nums1[k] = nums3[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
	}

	if i == m {
		for j < n {
			nums1[k] = nums2[j]
			j++
			k++
		}
	} else {
		for i < m {
			nums1[k] = nums3[i]
			i++
			k++
		}
	}
}

// 如果倒序放入数据到 nums1 中就不需要额外的空间
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			k--
			i--
		} else {
			nums1[k] = nums2[j]
			k--
			j--
		}
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
}
