package main

import (
	"errors"

	"golang.org/x/exp/constraints"
)

// A MaxHeap is a structure that stores elements in descending priority.
type MaxHeap[T constraints.Ordered] struct {
	data []T
}

// Peak returns the root node of the MaxHeap or an error if the heap is empty.
func (mh *MaxHeap[T]) Peak() (T, error) {
	if len(mh.data) == 0 {
		var emptyValue T
		return emptyValue, errors.New("heap is empty")
	}
	return mh.data[0], nil
}

// Insert inserts the given item into the MaxHeap and heapifys the heap into descending priority.
func (mh *MaxHeap[T]) Insert(item T) {
	mh.data = append(mh.data, item)
	mh.maxHeapifyUp(len(mh.data) - 1)
}

// Extract removes the root node from the MaxHeap and then heapifys the heap into descending priority.
// An error is returned if the heap is empty.
func (mh *MaxHeap[T]) Extract() (T, error) {
	if len(mh.data) == 0 {
		var emptyValue T
		return emptyValue, errors.New("heap is empty")
	}

	root := mh.data[0]

	if len(mh.data) == 1 {
		mh.data = []T{}
	} else {
		lastIdx := len(mh.data) - 1
		mh.data = append([]T{mh.data[lastIdx]}, mh.data[1:lastIdx]...)
		mh.maxHeapifyDown()
	}

	return root, nil
}

func (mh *MaxHeap[T]) maxHeapifyUp(currentIdx int) {
	for mh.data[parentIdx(currentIdx)] < mh.data[currentIdx] {
		mh.swapValues(currentIdx, parentIdx(currentIdx))
		currentIdx = parentIdx(currentIdx)
	}
}

func (mh *MaxHeap[T]) maxHeapifyDown() {
	parentIdx := 0
	for mh.parentSmallerThanChild(parentIdx) {
		lcIdx, rcIdx := leftChildIdx(parentIdx), rightChildIdx(parentIdx)
		if mh.data[lcIdx] > mh.data[rcIdx] {
			mh.swapValues(parentIdx, lcIdx)
			parentIdx = lcIdx
		} else {
			mh.swapValues(parentIdx, rcIdx)
			parentIdx = rcIdx
		}
	}
}

func (mh *MaxHeap[T]) swapValues(idx1, idx2 int) {
	mh.data[idx1], mh.data[idx2] = mh.data[idx2], mh.data[idx1]
}

func (mh *MaxHeap[T]) parentSmallerThanChild(parentIdx int) bool {
	parent := mh.data[parentIdx]
	lcIdx, rcIdx := leftChildIdx(parentIdx), rightChildIdx(parentIdx)
	if mh.childExists(lcIdx) && parent < mh.data[lcIdx] {
		return true
	}
	if mh.childExists(rcIdx) && parent < mh.data[rcIdx] {
		return true
	}
	return false
}

func (mh *MaxHeap[T]) childExists(idx int) bool {
	return idx <= len(mh.data)-1
}

func leftChildIdx(idx int) int {
	return idx*2 + 1
}

func rightChildIdx(idx int) int {
	return idx*2 + 2
}

func parentIdx(idx int) int {
	return (idx - 1) / 2
}
