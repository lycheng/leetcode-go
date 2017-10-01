package tree

import (
	"github.com/lycheng/leetcode-go/utils"
)

func mergeTrees(t1 *utils.TreeNode, t2 *utils.TreeNode) *utils.TreeNode {
	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	dst := &utils.TreeNode{}
	dst.Val = t1.Val + t2.Val
	if t1.Left != nil || t2.Left != nil {
		dst.Left = mergeTrees(t1.Left, t2.Left)
	}
	if t1.Right != nil || t2.Right != nil {
		dst.Right = mergeTrees(t1.Right, t2.Right)
	}
	return dst
}
