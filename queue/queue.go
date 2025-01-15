package queue

type Queue interface {
	InQueue(val int) bool
	DeQueue() int
}
