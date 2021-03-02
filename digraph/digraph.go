package digraph

import (
	"errors"
)

type digraph struct {
	vertices []string
	edges    []string
}

func (d digraph) IsGraphValid() error {
	for _,elem := range d.edges {
		if (2 < len(elem)) {
			return errors.New("edge cant have more than 2 vertices")
		}
		if (2 > len(elem)) {
			return errors.New("edge cant have less than 2 vertices")
		}
	}
	return nil
}

func (d digraph) ConstructIncidenceMatrix() [][]int {
	// for i, edge := range d.edges {
	// 	for j, vert := range d.vertices {
	// 	}
	// }
	return nil 
}