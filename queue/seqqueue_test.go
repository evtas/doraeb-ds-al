package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSeqQueue(t *testing.T) {
	sq := NewSeqQueue(10)
	assert.Equal(t, 0, sq.Q[9])
}

func TestSeqQueue_InQueue(t *testing.T) {
	sq := NewSeqQueue(1)
	ok := sq.InQueue(10)
	assert.Equal(t, true, ok)
	ok = sq.InQueue(11)
	assert.Equal(t, false, ok)
}

func TestSeqQueue_DeQueue(t *testing.T) {
	sq := NewSeqQueue(10)
	sq.InQueue(19)
	val := sq.DeQueue()
	assert.Equal(t, 19, val)
}
