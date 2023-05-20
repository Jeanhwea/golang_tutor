package nc0067

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0067_01(t *testing.T) {
	ans := uniquePaths(2, 2)
	assert.Equal(t, 2, ans)
}
