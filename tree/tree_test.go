package tree

import (
	"testing"

	"github.com/lycheng/leetcode-go/utils"
)

func Test_Traversal(t *testing.T) {

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
