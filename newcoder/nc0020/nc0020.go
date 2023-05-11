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
	count1 := merge(data, left, mid)
	count2 := merge(data, mid+1, right)

	temp := make([]int, len(data))
	copy(temp, data)
	i, j, k := left, mid+1, left
	for ; i <= mid && j <= right; k++ {
		if temp[i] < temp[j] {
			data[k] = temp[i]
			i++
		} else {
			data[k] = temp[j]
			j++
			ans += mid - i + 1
		}
	}

	for ; i <= mid; i++ {
		data[k] = temp[i]
		k++
	}

	for ; j <= right; j++ {
		data[k] = temp[j]
		k++
	}

	ans = (ans + count1 + count2) % 1000000007
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
