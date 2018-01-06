package tree

// https://leetcode.com/problems/two-sum-iv-input-is-a-bst/description/
func findTarget(root *TreeNode, k int) bool {
	set := make(map[int]bool)
	vals := inorderTraversal(root)
	for _, val := range vals {
		set[val] = true
	}

	for _, val := range vals {
		if k-val != val && set[k-val] {
			return true
		}
	}
	return false
}
