package quick

import (
	"testing"
)

func TestScratch01(t *testing.T) {
	a := [3]int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
			t.Log(a)
		}
		a[k] = 100 + v
	}
	t.Log(a)
}

func TestScratch02(t *testing.T) {
	a := []int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
			t.Log(a)
		}
		a[k] = 100 + v
	}
	t.Log(a)
}
