package graph_test

import (
	graph "github.com/ZombieHippie/Go-Concepts/DataStructures/graph"
	"testing"
)

func TestFindPath(t *testing.T) {
	var G graph.Digraph

	var first, last graph.V

	for i := 0; i < 5; i++ {
		var v1, v2 graph.V
		v1 = graph.V(i)
		v2 = graph.V(i + 1)

		if i == 0 {
			first = v1
		}
		last = v2

		G.AddEdge(v1, v2, float64(v1)/7+.1)
	}

	t.Log("Graph", G)
	t.Logf("First %d, Last %d", first, last)

	solved, solution := graph.FindPath(G, first, last)

	if !solved {
		t.Fatal("Failed by no solution.")
		t.FailNow()
	}

	VerifyPath(t, solution, first, last)

	t.Log("Solution result:", solution)
}

func VerifyPath(t *testing.T, solution []graph.EInterface, first, last graph.VInterface) {
	var lastEdge graph.EInterface
	lastIndiceOfSolution := len(solution)
	for index, edge := range solution {
		if index == 0 {
			if edge.From() != first {
				t.Fatal("First edge of solution set is not from start")
				t.FailNow()
			}
		} else if index <= lastIndiceOfSolution {
			if lastEdge.To() != edge.From() {
				t.Fatal("Solution set is disconnected path")
				t.FailNow()
			}
			if index == lastIndiceOfSolution {
				if edge.To() != last {
					t.Fatal("Last edge of solution set is not to goal")
					t.FailNow()
				}
			}
		}
		lastEdge = edge
	}

}
