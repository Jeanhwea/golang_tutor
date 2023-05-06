# 创建一个编写代码的文件

num="0001"

if [ $# -gt 0 ]; then
    num=$(printf %04d $1)
fi

# 源代码文件
file="$PWD/ncode/nc${num}/nc${num}.go"

# 检查文件是否存在
if [ -f "$file" ]; then
    echo "file already exists:\n  $file"
    exit 0
fi

# 新建包
pkg=$(dirname $file)

if [ ! -d "$pkg" ]; then
    mkdir -p $pkg
    echo "create package:\n  $pkg"
fi

# 创建 Model
model="${pkg}/model.go"
cat << EOF > $model
package nc${num}

type ListNode struct {
    Val  int
    Next *ListNode
}

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}
EOF

# 创建代码文件
cat << EOF > $file
package nc${num}
EOF

# 创建测试用例
test=${file/.go/_test.go}
cat << EOF > $test
package nc${num}

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func Test_NC${num}_01(t *testing.T) {
    assert.Equal(t, true, true)
}
EOF

echo "create file:\n  $model\n  $file\n  $test"
