package tree

import (
	"github.com/lycheng/leetcode-go/utils"
)

// https://leetcode.com/problems/binary-tree-inorder-traversal/
func inorderTraversal(root *utils.TreeNode) []int {
	result := make([]int, 0)

	if root == nil {
		return result
	}

	result = append(result, inorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)
	return result
}
