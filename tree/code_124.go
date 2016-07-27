package tree

import (
	"math"

	"github.com/lycheng/leetcode-go/utils"
)

func maxPathSum(root *utils.TreeNode) int {
	curMax := int(math.MinInt32)
	dfs(root, &curMax)
	return curMax
}

func dfs(root *utils.TreeNode, curMax *int) int {
	if root == nil {
		return 0
	}

	leftMax := dfs(root.Left, curMax)
	rightMax := dfs(root.Right, curMax)

	m := root.Val
	if leftMax > 0 {
		m += leftMax
	}
	if rightMax > 0 {
		m += rightMax
	}
	*curMax = maxInt(m, *curMax)
	return maxInt(root.Val, root.Val+leftMax, root.Val+rightMax)
}

func maxInt(ints ...int) int {
	m := ints[0]
	for _, i := range ints {
		if i > m {
			m = i
		}
	}
	return m
}
