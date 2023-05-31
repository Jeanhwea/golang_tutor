package nc0101

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0101_01(t *testing.T) {
	dlist := list.New()
	dlist.PushFront(3)
	assert.Equal(t, true, true)
}
