package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BiSearch(t *testing.T) {
	datas := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	val := 5
	index := BiSearch(datas, val)
	assert.Equal(t, 2, index)
}

func Test_BiSearchRecursion(t *testing.T) {
	datas := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	val := 5
	index := BiSearchRecursion(datas, val)
	assert.Equal(t, 2, index)
}

func Test_BiSearchFirstEqual(t *testing.T) {
	datas := []int{1, 3, 5, 5, 5, 7, 9, 11, 13, 15, 17, 19}
	val := 5
	index := BiSearchFirstEqual(datas, val)
	assert.Equal(t, 2, index)
	assert.NotEqual(t, 3, index)
}

func Test_BiSearchLastEqual(t *testing.T) {
	datas := []int{1, 3, 5, 5, 5, 7, 9, 11, 13, 15, 17, 19}
	val := 5
	index := BiSearchLastEqual(datas, val)
	assert.Equal(t, 4, index)
	assert.NotEqual(t, 3, index)
}

func Test_BiSearchFirstLargerEqual(t *testing.T) {
	datas := []int{1, 3, 5, 5, 5, 7, 9, 11, 13, 15, 17, 19}
	val := 6
	index := BiSearchFirstLarger(datas, val)
	assert.Equal(t, 5, index)
	assert.NotEqual(t, 4, index)
}

func Test_BiSearchLastLess(t *testing.T) {
	datas := []int{1, 3, 5, 5, 5, 7, 9, 11, 13, 15, 17, 19}
	val := 6
	index := BiSearchLastLess(datas, val)
	assert.Equal(t, 4, index)
	assert.NotEqual(t, 5, index)
}
