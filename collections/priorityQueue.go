package collections

type Item interface {
	Weight() float32
}

type PriorityQueue []*Item

func (pq *PriorityQueue) Len() int { return len(*pq) }

func (pq *PriorityQueue) Insert(e interface{}) {
	*pq = append(*pq, e.(*Item))
}

func (pq *PriorityQueue) IsEmpty() bool {
	return len(*pq) == 0
}

func (pq *PriorityQueue) DelMin() interface{} {
	di := -1
	var e *Item
	for i, v := range *pq {
		if di == -1 || (*v).Weight() < (*e).Weight() {
			di = i
			e = v
		}
	}
	if di == -1 {
		return nil
	}
	*pq = append((*pq)[:di], (*pq)[di+1:]...)
	return e
}

func (pq *PriorityQueue) Weight() float32 {
	var weight float32 = 0.0
	for _, edge := range *pq {
		weight += (*edge).Weight()
	}

	return weight
}
