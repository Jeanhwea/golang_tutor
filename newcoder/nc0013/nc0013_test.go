package nc0013

import (
	"testing"

	. "github.com/jeanhwea/golang_tutor/common/model"
	"github.com/stretchr/testify/assert"
)

func Test_NC0013_01(t *testing.T) {
	root := NewListNode([]int{1, 2, 1})
	assert.Equal(t, true, isPail(root))
}
