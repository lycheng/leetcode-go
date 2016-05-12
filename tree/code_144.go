package tree

import (
	"github.com/lycheng/leetcode-go/utils"
)

// https://leetcode.com/problems/binary-tree-preorder-traversal/
func preorderTraversal(root *utils.TreeNode) []int {
	result := make([]int, 0)

	if root == nil {
		return result
	}

	result = append(result, root.Val)
	result = append(result, preorderTraversal(root.Left)...)
	result = append(result, preorderTraversal(root.Right)...)
	return result
}
