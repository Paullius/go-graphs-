package collections

type Queue []int

func (q *Queue) Push(n int) {
	*q = append(*q, n)
}

func (q *Queue) Pop() (n int) {
	n = (*q)[0]
	*q = (*q)[1:]
	return
}

func (q *Queue) Len() int {
	return len(*q)
}
