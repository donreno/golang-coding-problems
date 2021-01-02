package structs

// GraphNode Representing a Graph node
type GraphNode struct {
	ID    int
	Value interface{}
	Nodes []*GraphNode
}

// MakeGraphNode creates a new GraphNode
func MakeGraphNode(id int, value interface{}) *GraphNode {
	return &GraphNode{
		ID:    id,
		Value: value,
	}
}

// GetNode gets an specific node by ID
func (g *GraphNode) GetNode(id int) *GraphNode {
	if g.ID == id {
		return g
	}

	for _, node := range g.Nodes {
		foundNode := node.GetNode(id)

		if foundNode != nil {
			return foundNode
		}
	}

	return nil
}

// AddNode Adds a node to a Graph Node
func (g *GraphNode) AddNode(node *GraphNode) {
	if len(g.Nodes) == 0 {
		g.Nodes = make([]*GraphNode, 0)
	}

	g.Nodes = append(g.Nodes, node)
}
