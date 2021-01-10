package graphsntrees

// FindAllPossibleBSTSequences finds all possible bst sequences
func FindAllPossibleBSTSequences(node *BinaryTreeNode) (possibleSequences [][]int) {
	possibleSequences = make([][]int, 0)

	if node == nil {
		possibleSequences = append(possibleSequences, []int{})
		return
	}

	prefix := []int{node.Val}

	leftSeq := FindAllPossibleBSTSequences(node.Left)
	rightSeq := FindAllPossibleBSTSequences(node.Right)

	for _, left := range leftSeq {
		for _, right := range rightSeq {
			weaved := weaveLists(left, right, [][]int{}, prefix)
			possibleSequences = append(possibleSequences, weaved...)
		}
	}

	return
}

func weaveLists(left, right []int, results [][]int, prefix []int) (weaved [][]int) {
	if len(left) == 0 || len(right) == 0 {
		result := append(prefix, left...)
		result = append(result, right...)
		weaved = append(results, result)
		return
	}

	weaved = weaveLists(left[1:], right, weaved, append(prefix, left[0]))
	weaved = weaveLists(left, right[1:], weaved, append(prefix, right[0]))

	return
}
