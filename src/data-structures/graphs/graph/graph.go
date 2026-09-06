package graph

import "errors"

var ErrInvalidIndex error = errors.New("function requires a valid index.")

type Graph interface {
	Directed() bool
	VertexCount() int
	Neighbors(vertex int, visit func(neighbor int, weight int64) (bool, error))
}
