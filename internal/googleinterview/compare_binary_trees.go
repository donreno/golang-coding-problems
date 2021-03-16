package googleinterview

// CompareTrees compares 2 binary trees
func CompareTrees(tree1, tree2 *TreeNode) bool {
	if tree1 == nil && tree2 == nil {
		return true
	} else if tree1 != nil && tree2 != nil {
		return tree1.Value == tree2.Value && CompareTrees(tree1.Left, tree2.Left) && CompareTrees(tree1.Right, tree2.Right)
	}

	return false
}
