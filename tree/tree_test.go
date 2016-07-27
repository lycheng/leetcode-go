package tree

import (
	"testing"

	"github.com/lycheng/leetcode-go/utils"
)

func TestTraversal(t *testing.T) {

	left := utils.TreeNode{2, nil, nil}
	right := utils.TreeNode{3, nil, nil}
	root := utils.TreeNode{1, &left, &right}

	preorder_result := preorderTraversal(&root)
	if preorder_result[0] != 1 {
		t.Error("error")
	}

	inorder_result := inorderTraversal(&root)
	if inorder_result[1] != 1 {
		t.Error("error")
	}

	postorder_result := postorderTraversal(&root)
	if postorder_result[2] != 1 {
		t.Error("error")
	}

	t.Log("success")
}

func TestCode124(t *testing.T) {
	left := utils.TreeNode{2, nil, nil}
	right := utils.TreeNode{3, nil, nil}
	root := utils.TreeNode{1, &left, &right}

	if maxPathSum(&root) != 6 {
		t.Error("error")
	}

	root.Left = &utils.TreeNode{-1, nil, nil}
	if maxPathSum(&root) != 4 {
		t.Error("error")
	}

	root.Right = &utils.TreeNode{-1, nil, nil}
	if maxPathSum(&root) != 1 {
		t.Error("error")
	}
	t.Log("success")
}
