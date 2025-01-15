package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewSeqStack(t *testing.T) {
	s := NewSeqStack(30)
	assert.Equal(t, int32(0), s.Data[29])
}

func TestStack_PushPop(t *testing.T) {
	s := NewSeqStack(10)
	s.Push(11)
	assert.Equal(t, int32(0), s.Top)
	assert.Equal(t, int32(11), s.Data[0])

	val := s.Pop()
	assert.Equal(t, int32(11), val)
}
