package googleinterview

// TreeNode represents a treeNode
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}

// MakeTreeNode makes a tree node
func MakeTreeNode(val int) *TreeNode {
	return &TreeNode{
		Value: val,
	}
}

// SwapTreeChildNodes swaps tree node childs
func SwapTreeChildNodes(node *TreeNode) {
	if node == nil {
		return
	}

	if node.Left != nil {
		SwapTreeChildNodes(node.Left)
	}

	if node.Right != nil {
		SwapTreeChildNodes(node.Right)
	}

	node.Left, node.Right = node.Right, node.Left
}
