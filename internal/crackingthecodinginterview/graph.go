package crackingthecodinginterview

// Graph Graph data struct
type Graph struct {
	nodeLookup map[int]*GraphNode
}

// GraphNode a Graph node
type GraphNode struct {
	Value  int
	Childs []*GraphNode
}

// MakeGraph makes a new Graph
func MakeGraph() *Graph {
	return &Graph{
		nodeLookup: make(map[int]*GraphNode),
	}
}

// GetNode gets a node by Value
func (g *Graph) GetNode(value int) *GraphNode {
	return g.nodeLookup[value]
}

// AddNode adds node to graph
func (g *Graph) AddNode(node *GraphNode) {
	if node == nil {
		return
	}

	g.nodeLookup[node.Value] = node
}

// AddEdge adds an edge to a given source node to link to destination node
func (g *Graph) AddEdge(source, destination int) {
	srcNode := g.nodeLookup[source]
	destNode := g.nodeLookup[destination]

	if srcNode != nil && destNode != nil {
		srcNode.Childs = append(srcNode.Childs, destNode)
	}
}

// HasPathDFS checks if there is a path to a node with DFS algorithm
func (g *Graph) HasPathDFS(source, destination int) bool {
	visited := make(map[int]bool)
	return g.hasPathDFS(source, destination, visited)
}

func (g *Graph) hasPathDFS(source, destination int, visited map[int]bool) bool {
	if visited[source] {
		return false
	}

	visited[source] = true

	if source == destination {
		return true
	}

	node := g.nodeLookup[source]

	for _, child := range node.Childs {
		if g.hasPathDFS(child.Value, destination, visited) {
			return true
		}
	}

	return false
}

// HasPathBFS has path to a node using BFS
func (g *Graph) HasPathBFS(source, destination int) bool {

	nextToVisit, current := make([]int, 0), 0
	nextToVisit = append(nextToVisit, source)

	for len(nextToVisit) > 0 {
		current, nextToVisit = nextToVisit[0], nextToVisit[1:]

		if current == destination {
			return true
		}

		childs := g.nodeLookup[current].Childs

		for _, child := range childs {
			nextToVisit = append(nextToVisit, child.Value)
		}
	}

	return false
}
