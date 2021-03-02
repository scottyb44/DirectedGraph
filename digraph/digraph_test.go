package digraph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func ConstructSimpleGraph() digraph {
	verts := []string{"A", "B", "C"}
	edges := []string{"AB", "BC", "CA"}
	return digraph{vertices: verts, edges: edges}
}

func TestDigraph_ValidGraph(t *testing.T) {
	digrph := ConstructSimpleGraph()

	err := digrph.IsGraphValid()

	require.Nil(t, err)
}

func TestDigraph_IncidenceMatrix_Exists(t *testing.T) {
	digrph := ConstructSimpleGraph()

	mat := digrph.ConstructIncidenceMatrix()
	
	require.NotNil(t, mat)
}

func TestDigraph_IncidenceMatrix_CorrectMatrix(t *testing.T) {
	digrph := ConstructSimpleGraph()

	mat := digrph.ConstructIncidenceMatrix()
	
	var expectedMatrix [][]int
	row1 := []int{1, 0, 1}
	row2 := []int{1, 1, 0}
	row3 := []int{0, 1, 1}
	expectedMatrix[0] = row1
	expectedMatrix[1] = row2
	expectedMatrix[2] = row3

	require.Equal(t, len(expectedMatrix), len(mat))
	// TODO SB: Need more test conditions
}