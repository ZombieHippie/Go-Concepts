package graph

// Interface borrowed from sort package
type VInterface interface{}
type EInterface interface {
	// get from vertice
	From() VInterface
	// get to vertice
	To() VInterface
	// get weight
	Weight() float64
}

// Directed graph
type DigraphInterface interface {
	// get a slice of the edges that point to v.
	To(VInterface) []EInterface
	// get a slice of the edges that point from v.
	From(VInterface) []EInterface
}

// Finds whatever path
func FindPath(G DigraphInterface, start, finish VInterface) (bool, []EInterface) {
	sol := make([]EInterface, 0, 2)
	return findPathR(G, start, finish, sol)
}

func findPathR(G DigraphInterface, start, finish VInterface, state []EInterface) (bool, []EInterface) {
	var lastVert VInterface

	if len(state) > 0 {
		lastVert = state[len(state)-1].To()
	} else {
		lastVert = start
	}

	if lastVert == finish {
		return true, state
	} else {
		nextEdges := G.From(lastVert)
		for _, edge := range nextEdges {
			test, solution := findPathR(G, start, finish, append(state, edge))
			if test {
				return true, solution
			}
		}
		return false, nil
	}
}
