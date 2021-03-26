package digraph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDigraph_AddNode(t *testing.T) {
	var graph DiGraph

	graph.AddNode(&Node{value: "A"})

	require.Equal(t, 1, len(graph.vertices))
}

func TestDigraph_AddEdge(t *testing.T) {
	var graph DiGraph

	graph.AddEdge(&Node{value: "A"}, &Node{value: "B"})

	require.Equal(t, 1, len(graph.edges))
}

// func TestDigraph_IncidenceMatrix_Exists(t *testing.T) {
// 	digrph := ConstructSimpleGraph()

// 	mat := digrph.ConstructIncidenceMatrix()
	
// 	require.NotNil(t, mat)
// }

// func TestDigraph_IncidenceMatrix_CorrectMatrix(t *testing.T) {
// 	digrph := ConstructSimpleGraph()

// 	mat := digrph.ConstructIncidenceMatrix()
	
// 	var expectedMatrix [][]int
// 	row1 := []int{1, 0, 1}
// 	row2 := []int{1, 1, 0}
// 	row3 := []int{0, 1, 1}
// 	expectedMatrix[0] = row1
// 	expectedMatrix[1] = row2
// 	expectedMatrix[2] = row3

// 	require.Equal(t, len(expectedMatrix), len(mat))
// 	// TODO SB: Need more test conditions
// }