package tree

import (
	"testing"

	"github.com/lycheng/leetcode-go/utils"
)

func Test_PreorderTraversal(t *testing.T) {

	left := utils.TreeNode{2, nil, nil}
	right := utils.TreeNode{3, nil, nil}

	root := utils.TreeNode{1, &left, &right}

	result := preorderTraversal(&root)

	if result[0] != 1 || result[2] != 3 {
		t.Error("error")
	}
	t.Log("Success")
}
