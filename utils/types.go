package utils

type Graph[T interface{}, W interface{}] struct {
	Vertices map[uint64]*Vertex[T, W]
}

type Vertex[T interface{}, W interface{}] struct {
	Value T
	Edges map[uint64]*Edge[T, W]
}

type Edge[T interface{}, W interface{}] struct {
	Weight W
	Vertex *Vertex[T, W]
}

func (g *Graph[T, W]) AddVertex(key uint64, value T) {
	g.Vertices[key] = &Vertex[T, W]{Value: value, Edges: map[uint64]*Edge[T, W]{}}
}

func (g *Graph[T, W]) AddEdge(key uint64, destKey uint64, weight W) {
	if _, ok := g.Vertices[key]; !ok {
		return
	}

	if _, ok := g.Vertices[destKey]; !ok {
		return
	}

	g.Vertices[key].Edges[destKey] = &Edge[T, W]{Weight: weight, Vertex: g.Vertices[destKey]}
}

func (g *Graph[T, W]) Neighbors(key uint64) []T {
	result := []T{}

	for _, edge := range g.Vertices[key].Edges {
		result = append(result, edge.Vertex.Value)
	}

	return result
}
