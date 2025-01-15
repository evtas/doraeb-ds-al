package queue

import "container/list"

type SeqQueue struct {
	Q    []int
	N    int
	Head int
	Tail int
}

func NewSeqQueue(n int) *SeqQueue {
	list.New()
	return &SeqQueue{
		// 会浪费一个空间，所以创建时需要n+1
		Q:    make([]int, n+1),
		N:    n + 1,
		Head: 0,
		Tail: 0,
	}
}

func (sq *SeqQueue) InQueue(val int) bool {
	if (sq.Tail+1)%sq.N == sq.Head {
		return false
	}
	sq.Q[sq.Tail] = val
	sq.Tail += 1
	return true
}

func (sq *SeqQueue) DeQueue() int {
	val := sq.Q[sq.Head]
	sq.Head += 1
	sq.Head = sq.Head % sq.N
	return val
}
