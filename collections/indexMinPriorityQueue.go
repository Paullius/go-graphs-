package collections

type IndexMinItem struct {
	value float32 // The value of the item; arbitrary.

	index int // The index of the item in the heap.
}

type IndexMinPriorityQueue []IndexMinItem

func (pq IndexMinPriorityQueue) Len() int { return len(pq) }

func (pq *IndexMinPriorityQueue) Insert(i int, value float32) {
	item := IndexMinItem{
		index: i,
		value: value}

	*pq = append(*pq, item)
}

func (pq *IndexMinPriorityQueue) DelMin() int {
	di := -1
	mi := -1
	var dv float32
	for i, v := range *pq {
		if di == -1 || v.value < dv {
			di = i
			mi = v.index
			dv = v.value
		}
	}

	*pq = append((*pq)[:di], (*pq)[di+1:]...)

	return mi
}

func (pq *IndexMinPriorityQueue) Contains(i int) bool {

	for _, v := range *pq {
		if v.index == i {
			return true
		}
	}
	return false
}

func (pq *IndexMinPriorityQueue) Change(i int, value float32) bool {

	for _, v := range *pq {
		if v.index == i {
			v.value = value
			return true
		}
	}
	return false
}

func (pq *IndexMinPriorityQueue) IsEmpty() bool {
	return len(*pq) == 0
}
