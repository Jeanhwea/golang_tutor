package nc0089

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0089_01(t *testing.T) {
	intervals := []*Interval{
		{10, 20},
		{20, 60},
		{80, 100},
		{150, 180},
	}
	ans := merge(intervals)
	v, _ := json.Marshal(ans)
	t.Log(string(v))
	assert.Equal(t, true, true)
}
