package nc0020

func InversePairs(data []int) (ans int) {
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
