package weighted

type Vertex struct {
	index  int
	weight float32
}

func (i *Vertex) Weight() float32 {
	return i.weight
}

func (i *Vertex) Index() int {
	return i.index
}

func (i *Vertex) SetWeight(weight float32) {
	i.weight = weight
}

func NewVertex(intdex int, weight float32) Vertex {
	return Vertex{intdex, weight}
}
