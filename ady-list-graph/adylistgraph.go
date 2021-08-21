package adylistgraph

type adygraph struct {
	list [][]*Vertex
}

type Vertex struct {
	value interface{}
}

func newAdyGraph() *adygraph {
	return &adygraph{make([][]*Vertex, 0)}
}
