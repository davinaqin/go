package queue

type Queue []int

func (q *Queue)Push(i int)  {
	*q = append(*q, i)
}

func (q *Queue)Pop() int {
	i := (*q)[0]
	*q = (*q)[1:]
	return i
}

func (q *Queue) IsEmpty() bool {
	if len(*q) == 0 {
		return false
	}
	return true
}


