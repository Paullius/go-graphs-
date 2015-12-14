package weighted

type PriorityQueue []*Edge

func (pq *PriorityQueue) Len() int { return len(*pq) }

func (pq *PriorityQueue) Insert(e *Edge) {
	*pq = append(*pq, e)
}

func (pq *PriorityQueue) IsEmpty() bool {
	return len(*pq) == 0
}

func (pq *PriorityQueue) DelMin() *Edge {
	di := -1
	var e *Edge
	for i, v := range *pq {
		if di == -1 || v.weight < e.weight {
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
		weight += edge.weight
	}

	return weight
}
