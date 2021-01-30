package sorting

type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

// MakeBinaryTree makes binary tree
func MakeBinaryTree() *BinaryTreeNode {
	return &BinaryTreeNode{Val: -1}
}

// Insert insterts element into binary tree
func (b *BinaryTreeNode) Insert(val int) {
	if b.Val == -1 {
		b.Val = val
		return
	}

	if val <= b.Val {
		if b.Left == nil {
			b.Left = MakeBinaryTree()
		}

		b.Left.Insert(val)
	} else {
		if b.Right == nil {
			b.Right = MakeBinaryTree()
		}

		b.Right.Insert(val)
	}
}

// RankFromStream gets rank from stream
func RankFromStream(stream *BinaryTreeNode, elem int) int {
	if stream == nil {
		return 0
	}

	// If the item is equal to val or greater than, then I want to look to the left
	if stream.Val >= elem {
		if stream.Left != nil {
			return RankFromStream(stream.Left, elem)
		}

		return 0
	}

	left, right := 0, 0

	if stream.Left != nil {
		left = RankFromStream(stream.Left, elem)
	}

	if stream.Right != nil {
		right = RankFromStream(stream.Right, elem)
	}

	return left + right + 1
}
