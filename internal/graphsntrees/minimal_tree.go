package graphsntrees

// BinaryTreeNode binary search tree node
type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

// FillMinimalBinaryTree fills binary search tree as balanced as possible
func FillMinimalBinaryTree(elems []int) *BinaryTreeNode {
	root := new(BinaryTreeNode)

	fillMinimalBinaryTree(elems, root)

	return root
}

func fillMinimalBinaryTree(elems []int, root *BinaryTreeNode) {
	lenElems := len(elems)

	if lenElems <= 3 {
		switch lenElems {
		case 1:
			root.Val = elems[0]
		case 2:
			root.Left = &BinaryTreeNode{Val: elems[0]}
			root.Val = elems[1]
		case 3:
			root.Left = &BinaryTreeNode{Val: elems[0]}
			root.Val = elems[1]
			root.Right = &BinaryTreeNode{Val: elems[2]}
		}

		return
	}

	midPoint := lenElems / 2

	root.Val = elems[midPoint]
	root.Left = new(BinaryTreeNode)
	root.Right = new(BinaryTreeNode)

	fillMinimalBinaryTree(elems[:midPoint], root.Left)
	fillMinimalBinaryTree(elems[midPoint+1:], root.Right)
}
