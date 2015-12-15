package collections

type Stack []interface{}

func (q *Stack) Push(n interface{}) {
	*q = append(*q, n)
}

func (q *Stack) Pop() (n interface{}) {
	x := q.Len() - 1
	n = (*q)[x]
	*q = (*q)[:x]
	return
}
func (q *Stack) Len() int {
	return len(*q)
}

func (st *Stack) ConvertToInt() []int {
	b := make([]int, len(*st))
	for i, s := range *st {
		b[i] = s.(int)
	}
	return b
}
