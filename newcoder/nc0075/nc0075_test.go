package nc0075

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0075_01(t *testing.T) {
	str1, str2 := "nowcoder", "new"
	ans := editDistance(str1, str2)
	assert.Equal(t, 6, ans)
}


func Test_NC0075_02(t *testing.T) {
	str1, str2 := "lrbb","mqbhcda"
	ans := editDistance(str1, str2)
	assert.Equal(t, 6, ans)
}
