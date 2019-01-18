package tree

import (
	"reflect"
	"testing"
)

func TestTraversal(t *testing.T) {

	left := TreeNode{2, nil, nil}
	right := TreeNode{3, nil, nil}
	root := TreeNode{1, &left, &right}

	preorderRv := preorderTraversal(&root)
	if preorderRv[0] != 1 {
		t.Error("error")
	}

	inorderRv := inorderTraversal(&root)
	if inorderRv[1] != 1 {
		t.Error("error")
	}

	postorderRv := postorderTraversal(&root)
	if postorderRv[2] != 1 {
		t.Error("error")
	}

	t.Log("success")
}

func TestCode124(t *testing.T) {
	left := TreeNode{2, nil, nil}
	right := TreeNode{3, nil, nil}
	root := TreeNode{1, &left, &right}

	if maxPathSum(&root) != 6 {
		t.Error("error")
	}

	root.Left = &TreeNode{-1, nil, nil}
	if maxPathSum(&root) != 4 {
		t.Error("error")
	}

	root.Right = &TreeNode{-1, nil, nil}
	if maxPathSum(&root) != 1 {
		t.Error("error")
	}
	t.Log("success")
}

func TestCode617(t *testing.T) {
	t1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	dst := mergeTrees(t1, nil)
	if t1 != dst {
		t.Fatal("error with t2 = nil")
	}

	t2 := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 3,
		},
	}

	dst = mergeTrees(t1, t2)
	if dst.Val != 2 || dst.Left.Val != 2 || dst.Right.Val != 6 {
		t.Error("error with t2 = nil")
	}
	t.Log("Success")
}

func TestCode637(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
		},
	}
	expected := []float64{1, 2.5, 4}
	if ret := averageOfLevels(root); !reflect.DeepEqual(ret, expected) {
		t.Fatal("expect", expected, "and get", ret)
	}
	t.Log("success")
}

func TestCode653(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	k := 9
	expected := true
	if ret := findTarget(root, k); ret != expected {
		t.Fatal("expect", expected, "and get", ret)
	}

	k = 28
	expected = false
	if ret := findTarget(root, k); ret != expected {
		t.Fatal("expect", expected, "and get", ret)
	}

	root = &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	k = 4
	expected = true
	if ret := findTarget(root, k); ret != expected {
		t.Fatal("expect", expected, "and get", ret)
	}

	k = 1
	expected = false
	if ret := findTarget(root, k); ret != expected {
		t.Fatal("expect", expected, "and get", ret)
	}
	t.Log("success")
}
