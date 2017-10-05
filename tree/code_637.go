package tree

// https://leetcode.com/problems/average-of-levels-in-binary-tree/description/
func averageOfLevels(root *TreeNode) []float64 {
	var ret []float64
	if root == nil {
		return ret
	}
	nodes := make([]*TreeNode, 0)
	nodes = append(nodes, root)

	for len(nodes) > 0 {
		var f bool
		var curSum float64
		var node *TreeNode
		var nNodes []*TreeNode
		cnt := float64(len(nodes))
		for len(nodes) > 0 {
			f = true
			node, nodes = nodes[0], nodes[1:]
			curSum += float64(node.Val)
			if node.Left != nil {
				nNodes = append(nNodes, node.Left)
			}
			if node.Right != nil {
				nNodes = append(nNodes, node.Right)
			}
		}
		if f {
			ret = append(ret, curSum/cnt)
		}
		nodes = nNodes
	}
	return ret
}
