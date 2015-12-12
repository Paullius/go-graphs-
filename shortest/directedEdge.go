package shortest

type DirectedEdge struct {
	v, w   int
	weight float32
}

func NewEdge(v int, w int, weight float32) DirectedEdge {
	e := DirectedEdge{v, w, weight}

	return e
}

func (e *DirectedEdge) Weight() float32 {
	return e.weight
}

func (e *DirectedEdge) From() int {
	return e.v
}

func (e *DirectedEdge) To() int {
	return e.w
}

func (e *DirectedEdge) Initialized() bool {
	return e.v != 0 || e.w != 0
}
