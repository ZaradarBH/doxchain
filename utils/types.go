package types

type Graph[T interface{}, W interface{}] struct {
	Vertices map[int]*Vertex[T, W]
}

type Vertex[T interface{}, W interface{}] struct {
	Value T
	Edges map[int]*Edge[T, W]
}

type Edge[T interface{}, W interface{}] struct {
	Weight W
	Vertex *Vertex[T, W]
}

func (this *Graph[T, W]) AddVertex(key int, value T) {
	this.Vertices[key] = &Vertex[T, W]{Value: value, Edges: map[int]*Edge[T, W]{}}
}

func (this *Graph[T, W]) AddEdge(srcKey int, destKey int, weight W) {
	// check if src & dest exist
	if _, ok := this.Vertices[srcKey]; !ok {
		return
	}

	if _, ok := this.Vertices[destKey]; !ok {
		return
	}

	this.Vertices[srcKey].Edges[destKey] = &Edge[T, W]{Weight: weight, Vertex: this.Vertices[destKey]}
}

func (this *Graph[T, W]) Neighbors(srcKey int) []T {
	result := []T{}

	for _, edge := range this.Vertices[srcKey].Edges {
		result = append(result, edge.Vertex.Value)
	}

	return result
}
