package recursionndynamic

import (
	s "golang-coding-problems/internal/structs"
)

const maxBoxSurface = 9999

// Box represents a box
type Box struct {
	Width  int
	Depth  int
	Height int
}

// Surface gets box surface
func (b Box) Surface() int {
	return b.Width * b.Depth
}

// FitsOnTopOf checks if a box fits on top of another
func (b *Box) FitsOnTopOf(ob Box) bool {
	return b.Width < ob.Width && b.Depth < ob.Depth
}

// StackOfBoxes Returns tallest possible stack of boxes
func StackOfBoxes(boxes []*Box) int {
	height := 0

	if len(boxes) == 0 {
		return height
	}

	//build surface matrix, so we can group elements with same surface on a list
	surfaceMatrix := make(map[int]*s.LinkedList)
	maxSurface := 0
	minSurface := maxBoxSurface
	for _, box := range boxes {
		if box.Surface() > maxSurface {
			maxSurface = box.Surface()
		}

		if box.Surface() < minSurface {
			minSurface = box.Surface()
		}

		addBoxToSurfaceMatrix(box, surfaceMatrix)
	}

	var newBottom *Box

	for i := maxSurface; i >= minSurface; i-- {
		if _, present := surfaceMatrix[i]; !present {
			continue
		}

		tallestFittingBox := getBestFitForCurrentBottom(surfaceMatrix[i], newBottom)

		if tallestFittingBox != nil {
			height += tallestFittingBox.Height
			newBottom = tallestFittingBox
		}
	}

	return height
}

func addBoxToSurfaceMatrix(box *Box, surfaceMatrix map[int]*s.LinkedList) {
	surface := box.Surface()
	if surfaceMatrix[surface] == nil {
		surfaceMatrix[surface] = new(s.LinkedList)
	}

	surfaceMatrix[surface].Add(box)
}

func getBestFitForCurrentBottom(boxesList *s.LinkedList, bottomBox *Box) *Box {

	var tallestFittingBox *Box = nil
	for _, elem := range boxesList.ToSlice() {
		box := elem.(*Box)

		if (bottomBox == nil || box.FitsOnTopOf(*bottomBox)) &&
			(tallestFittingBox == nil || box.Height > tallestFittingBox.Height) {
			tallestFittingBox = box
		}
	}

	return tallestFittingBox
}
