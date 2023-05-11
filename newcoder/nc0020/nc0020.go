package nc0020

// 逆序对数量
func InversePairs(data []int) (ans int) {
	return merge(data, 0, len(data)-1)
}

func merge(data []int, left, right int) (ans int) {
	if left >= right {
		return
	}
	mid := left + (right-left)/2
	cnt1, cnt2 := merge(data, left, mid), merge(data, mid+1, right)

	temp := append([]int{}, data[left:right+1]...)
	i, j, k := 0, mid-left+1, left
	for ; i <= mid-left && j <= right-left; k++ {
		if temp[i] < temp[j] {
			data[k] = temp[i]
			i++
		} else {
			data[k] = temp[j]
			j++
			ans += mid - left - i + 1
		}
	}

	for ; i <= mid-left; i++ {
		data[k] = temp[i]
		k++
	}

	for ; j <= right-left; j++ {
		data[k] = temp[j]
		k++
	}

	ans = (ans + cnt1 + cnt2) % 1000000007
	return
}

func InversePairs2(data []int) (ans int) {
	n := len(data)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if data[i] > data[j] {
				ans++
			}
		}
	}
	return
}
