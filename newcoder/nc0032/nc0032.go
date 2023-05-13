package nc0032

import (
    . "github.com/jeanhwea/golang_tutor/common/model"
)

func mergeTrees(t1, t2 *TreeNode ) (ans *TreeNode) {
    return merge(t1, t2)
}

func merge(root1, root2 *TreeNode) (ans *TreeNode) {
    if root1 == nil && root2 == nil {
        return
    }

    if root1 != nil && root2 != nil {
        ans = &TreeNode{
            Val: root1.Val + root2.Val,
            Left: merge(root1.Left, root2.Left),
            Right: merge(root2.Right, root2.Right),
        }
        return
    }

    if root1 != nil {
        ans = &TreeNode{
            Val: root1.Val,
            Left: merge(root1.Left, nil),
            Right: merge(root1.Right, nil),
        }
    }

    if root2 != nil {
        ans = &TreeNode{
            Val: root2.Val,
            Left: merge(nil, root2.Left),
            Right: merge(nil, root2.Right),
        }
    }

    return
}