package lc0322

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LC0322_01(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	ans := coinChange(coins, amount)
	assert.Equal(t, 3, ans)
}

func Test_LC0322_02(t *testing.T) {
	coins := []int{186, 419, 83, 408}
	amount := 6249
	ans := coinChange(coins, amount)
	assert.Equal(t, 20, ans)
}
