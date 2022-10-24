package main

import (
	"errors"
)

// A MaxHeap is a structure that stores elements in descending priority.
type MaxHeap struct {
	data []int
}

// Peak returns the root node of the MaxHeap or an error if the heap is empty.
func (mh MaxHeap) Peak() (int, error) {
	if len(mh.data) == 0 {
		return -1, errors.New("heap is empty")
	}
	return mh.data[0], nil
}

// Insert inserts the given item into the MaxHeap and heapifys the heap into descending priority.
func (mh *MaxHeap) Insert(item int) {
	mh.data = append(mh.data, item)
	mh.maxHeapifyUp(len(mh.data) - 1)
}

// Extract removes the root node from the MaxHeap and then heapifys the heap into descending priority.
// An error is returned if the heap is empty.
func (mh *MaxHeap) Extract() (int, error) {
	if len(mh.data) == 0 {
		return -1, errors.New("heap is empty")
	}

	root := mh.data[0]

	if len(mh.data) == 1 {
		mh.data = []int{}
	} else {
		lastIdx := len(mh.data) - 1
		mh.data = append([]int{mh.data[lastIdx]}, mh.data[1:lastIdx]...)
		mh.maxHeapifyDown()
	}

	return root, nil
}

func (mh *MaxHeap) maxHeapifyUp(currentIdx int) {
	for mh.data[parentIdx(currentIdx)] < mh.data[currentIdx] {
		mh.swapValues(currentIdx, parentIdx(currentIdx))
		currentIdx = parentIdx(currentIdx)
	}
}

func (mh *MaxHeap) maxHeapifyDown() {
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

func (mh *MaxHeap) swapValues(idx1, idx2 int) {
	mh.data[idx1], mh.data[idx2] = mh.data[idx2], mh.data[idx1]
}

func (mh *MaxHeap) parentSmallerThanChild(parentIdx int) bool {
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

func (mh *MaxHeap) childExists(idx int) bool {
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
