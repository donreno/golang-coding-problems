package recursionndynamic

import (
	s "golang-coding-problems/internal/structs"
)

// StackOfBoxesRecursive stack of boxes recursive
func StackOfBoxesRecursive(boxes []*Box) int {
	boxes = sortBoxes(boxes)
	stackMap := make([]int, len(boxes))

	return stackOfBoxes(boxes, nil, 0, stackMap)
}

func stackOfBoxes(boxes []*Box, bottom *Box, offset int, stackMap []int) int {

	if offset >= len(boxes) {
		return 0
	}

	newBottom := boxes[offset]
	heightWithBottom := 0

	if bottom == nil || newBottom.FitsOnTopOf(*bottom) {
		if stackMap[offset] == 0 {
			stackMap[offset] = stackOfBoxes(boxes, newBottom, offset+1, stackMap)
			stackMap[offset] += newBottom.Height
		}
		heightWithBottom = stackMap[offset]
	}

	heightWithoutBottom := stackOfBoxes(boxes, bottom, offset+1, stackMap)

	if heightWithBottom < heightWithoutBottom {
		return heightWithoutBottom
	}

	return heightWithBottom
}

func sortBoxes(boxes []*Box) []*Box {
	heightMap := make(map[int]*s.LinkedList)
	maxheight := 0

	for _, box := range boxes {
		if box.Height > maxheight {
			maxheight = box.Height
		}

		addBoxToHeightMap(box, heightMap)
	}

	for i, h := 0, maxheight; h > 0; h-- {
		if _, present := heightMap[h]; !present {
			continue
		}

		for _, item := range heightMap[h].ToSlice() {
			boxes[i] = item.(*Box)
			i++
		}
	}

	return boxes
}

func addBoxToHeightMap(box *Box, heightMap map[int]*s.LinkedList) {
	height := box.Height
	if heightMap[height] == nil {
		heightMap[height] = new(s.LinkedList)
	}

	heightMap[height].Add(box)
}
