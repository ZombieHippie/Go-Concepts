package graph

// Interface borrowed from sort package
type V int
type E struct {
	from, to V
	weight   float64
}

func (e E) From() VInterface { return e.from }
func (e E) To() VInterface   { return e.to }
func (e E) Weight() float64  { return e.weight }

// Directed graph
type Digraph struct {
	edges []E
}

func (d Digraph) To(to VInterface) []EInterface {
	toEdges := make([]EInterface, 0, 2)
	for _, edge := range d.edges {
		if edge.To() == to {
			toEdges = append(toEdges, edge)
		}
	}
	return toEdges
}

func (d Digraph) From(from VInterface) []EInterface {
	fromEdges := make([]EInterface, 0, 2)
	for _, edge := range d.edges {
		if edge.From() == from {
			fromEdges = append(fromEdges, edge)
		}
	}
	return fromEdges
}

func (d *Digraph) AddEdge(from, to V, weight float64) {
	d.edges = append(d.edges, E{from, to, weight})
}
