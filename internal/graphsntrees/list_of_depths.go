package graphsntrees

// LinkedNode node for linked list node
type LinkedNode struct {
	Val  int
	Next *LinkedNode
}

// ListOfDepths returns a list with all the depths
func ListOfDepths(root *BinaryTreeNode) map[int]*LinkedNode {
	listOfDepths := make(map[int]*LinkedNode)

	fillListOfDepths(root, listOfDepths, 0)

	return listOfDepths
}

func fillListOfDepths(node *BinaryTreeNode, listOfDepths map[int]*LinkedNode, currentLevel int) {
	if node == nil {
		return
	}

	listOfDepths[currentLevel] = &LinkedNode{
		Val:  node.Val,
		Next: listOfDepths[currentLevel],
	}

	fillListOfDepths(node.Left, listOfDepths, currentLevel+1)
	fillListOfDepths(node.Right, listOfDepths, currentLevel+1)
}
