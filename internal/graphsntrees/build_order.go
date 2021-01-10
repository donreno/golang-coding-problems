package graphsntrees

// BuildOrder returns build order
func BuildOrder(projects []rune, dependencyGraph map[rune][]rune) map[int]rune {
	built := make(map[rune]bool)
	buildOrder := make(map[int]rune)
	buildIndex := 0

	for _, project := range projects {
		deps := dependencyGraph[project]

		if built[project] {
			continue
		}

		for _, dep := range deps {
			if !built[dep] {
				buildOrder[buildIndex] = dep
				built[dep] = true
				buildIndex++
			}
		}

		buildOrder[buildIndex] = project
		built[project] = true
		buildIndex++
	}

	return buildOrder
}
