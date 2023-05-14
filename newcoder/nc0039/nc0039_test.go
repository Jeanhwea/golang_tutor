package nc0039

import (
	"encoding/json"
	"testing"

	. "github.com/jeanhwea/golang_tutor/common/model"
	"github.com/stretchr/testify/assert"
)

func Test_NC0039_01(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 6},
		},
	}
	ans := Serialize(root)
	assert.Equal(t, "*1,*2,*4,#,#,*5,#,#,*3,#,*6,#,#", ans)
}

func TestNc003901(t *testing.T) {
	s := "*1,*2,*4,#,#,*5,#,#,*3,#,*6,#,#"
	root := Deserialize(s)
	str, _ := json.MarshalIndent(root, "", "  ")
	t.Log(string(str))
}
