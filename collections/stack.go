package collections

type Stack []int

func (q *Stack) Push(n int) {
	*q = append(*q, n)
}

func (q *Stack) Pop() (n int) {
	x := q.Len() - 1
	n = (*q)[x]
	*q = (*q)[:x]
	return
}
func (q *Stack) Len() int {
	return len(*q)
}
