package graphsntrees

// CheckT2IsSubtreeOfT1 checks if t2 is subtree of t1
func CheckT2IsSubtreeOfT1(t1, t2 *BinaryTreeNode) bool {
	if t2 == nil {
		return true
	}

	return subTree(t1, t2)
}

func subTree(t1, t2 *BinaryTreeNode) bool {
	if t1 == nil {
		return false
	} else if t1.Val == t2.Val && matchTree(t1, t2) {
		return true
	}

	return subTree(t1.Left, t2) || subTree(t1.Right, t2)
}

func matchTree(t1, t2 *BinaryTreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	} else if t1 == nil || t2 == nil {
		return false
	} else if t1.Val != t2.Val {
		return false
	}

	return matchTree(t1.Left, t2.Left) && matchTree(t1.Right, t2.Right)
}
