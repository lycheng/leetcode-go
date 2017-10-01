package tree

import (
	"github.com/lycheng/leetcode-go/utils"
)

// https://leetcode.com/problems/binary-tree-postorder-traversal/
func postorderTraversal(root *utils.TreeNode) []int {
	result := make([]int, 0)

	if root == nil {
		return result
	}

	result = append(result, postorderTraversal(root.Left)...)
	result = append(result, postorderTraversal(root.Right)...)
	result = append(result, root.Val)
	return result
}
