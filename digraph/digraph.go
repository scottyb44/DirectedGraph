package digraph

import (
	"fmt"
)

type Node struct {
	value string
}

type DiGraph struct {
	vertices []*Node
	edges    map[Node][]*Node
}

func (d *DiGraph) AddNode(node *Node) {
	d.vertices = append(d.vertices, node)
}

func (d *DiGraph) AddEdge(vert1 *Node, vert2 *Node) {
	if d.edges == nil {
		d.edges = make(map[Node][]*Node)
	}
	d.edges[*vert1] = append(d.edges[*vert1], vert2)
}

func (d *DiGraph) PrintGraph() {
	strGraph := ""
	for _, elem := range d.vertices {
		strGraph += elem.value + "-> "

		edges, ok := d.edges[*elem]
		if !ok {
			continue
		}
		for _, edge := range edges {
			strGraph += edge.value + " "
		}
		strGraph += "\n"
	}
	fmt.Println(strGraph)
}

// TODO SB: Do this and make an adjacency matrix
func (d DiGraph) ConstructIncidenceMatrix() [][]int {
	// for i, edge := range d.edges {
	// 	for j, vert := range d.vertices {
	// 	}
	// }
	return nil
}