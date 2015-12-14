package weighted

type Edge struct {
	v, w   int
	weight float32
}

func NewEdge(v int, w int, weight float32) Edge {
	e := Edge{v, w, weight}

	return e
}

func (e *Edge) OtherVertex(vertex int) int {
	if vertex == e.v {
		return e.w
	} else if vertex == e.w {
		return e.v
	}

	return -1
}

func (e *Edge) Weight() float32 {
	return e.weight
}

func (e *Edge) From() int {
	return e.v
}

func (e *Edge) To() int {
	return e.w
}

func (e *Edge) CompareTo(edge *Edge) int {
	if e.Weight() < edge.Weight() {
		return -1
	} else if e.Weight() > edge.Weight() {
		return +1
	} else {
		return 0
	}
}

func (e *Edge) Initialized() bool {
	return e.v != 0 || e.w != 0
}
