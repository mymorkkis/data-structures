package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeapPeakReturnsRootNode(t *testing.T) {
	mh := MaxHeap[int]{[]int{15, 10}}

	root, _ := mh.Peak()

	assert.Equal(t, 15, root)
}

func TestMaxHeapPeakErrorsWhenHeapEmpty(t *testing.T) {
	mh := MaxHeap[int]{}

	root, err := mh.Peak()

	assert.Empty(t, root)
	assert.Equal(t, errors.New("heap is empty"), err)
}

func TestMaxHeapInsert(t *testing.T) {
	mh := MaxHeap[int]{[]int{15, 10}}

	mh.Insert(5)

	assert.Equal(t, mh.data, []int{15, 10, 5})

	mh.Insert(20)

	assert.Equal(t, mh.data, []int{20, 15, 5, 10})

	mh.Insert(17)

	assert.Equal(t, mh.data, []int{20, 17, 5, 10, 15})

	mh.Insert(25)

	assert.Equal(t, mh.data, []int{25, 17, 20, 10, 15, 5})
}

func TestMaxHeapExtractRemovesRootNode(t *testing.T) {
	mh := MaxHeap[int]{[]int{20, 15, 10, 5, 12}}

	previous_root, _ := mh.Extract()

	assert.Equal(t, 20, previous_root)
	assert.Equal(t, []int{15, 12, 10, 5}, mh.data)

	previous_root, _ = mh.Extract()

	assert.Equal(t, 15, previous_root)
	assert.Equal(t, []int{12, 5, 10}, mh.data)

	previous_root, _ = mh.Extract()

	assert.Equal(t, 12, previous_root)
	assert.Equal(t, []int{10, 5}, mh.data)

	previous_root, _ = mh.Extract()
	assert.Equal(t, 10, previous_root)
	assert.Equal(t, []int{5}, mh.data)

	previous_root, _ = mh.Extract()
	assert.Equal(t, 5, previous_root)
	assert.Equal(t, []int{}, mh.data)
}

func TestMaxHeapExtractErrorsWhenHeapEmpty(t *testing.T) {
	mh := MaxHeap[int]{}

	previousRoot, err := mh.Extract()

	assert.Empty(t, previousRoot)
	assert.Equal(t, errors.New("heap is empty"), err)
}

func TestMaxHeapCanStoreStrings(t *testing.T) {
	mh := MaxHeap[string]{}
	mh.Insert("second")
	mh.Insert("first")

	root, _ := mh.Peak()
	assert.Equal(t, "second", root)

	previousRoot, _ := mh.Extract()
	assert.Equal(t, "second", previousRoot)

	previousRoot, _ = mh.Extract()
	assert.Equal(t, "first", previousRoot)

	emptyValue, err := mh.Extract()
	assert.Empty(t, emptyValue)
	assert.NotNil(t, err)
}

func TestMaxHeapCanStoreFloats(t *testing.T) {
	mh := MaxHeap[float64]{}
	mh.Insert(2.4)
	mh.Insert(1.2)

	root, _ := mh.Peak()
	assert.Equal(t, 2.4, root)

	previousRoot, _ := mh.Extract()
	assert.Equal(t, 2.4, previousRoot)

	previousRoot, _ = mh.Extract()
	assert.Equal(t, 1.2, previousRoot)

	emptyValue, err := mh.Extract()
	assert.Empty(t, emptyValue)
	assert.NotNil(t, err)
}
