package main

import (
	"errors"
)

type MaxHeap struct {
	data []int
}

func (heap MaxHeap) Peak() (int, error) {
	if len(heap.data) == 0 {
		return -1, errors.New("heap is empty")
	}
	return heap.data[0], nil
}

func (heap *MaxHeap) Insert(key int) {
	heap.data = append(heap.data, key)
	heap.maxHeapifyUp(len(heap.data) - 1)
}

func (heap *MaxHeap) Extract() (int, error) {
	if len(heap.data) == 0 {
		return -1, errors.New("heap is empty")
	}

	head := heap.data[0]

	if len(heap.data) == 1 {
		heap.data = []int{}
	} else {
		lastIdx := len(heap.data) - 1
		heap.data = append([]int{heap.data[lastIdx]}, heap.data[1:lastIdx]...)
		heap.maxHeapifyDown()
	}

	return head, nil
}

func (heap *MaxHeap) maxHeapifyUp(currentIdx int) {
	for heap.data[parentIdx(currentIdx)] < heap.data[currentIdx] {
		heap.swapValues(currentIdx, parentIdx(currentIdx))
		currentIdx = parentIdx(currentIdx)
	}
}

func (heap *MaxHeap) maxHeapifyDown() {
	parentIdx := 0
	for heap.parentSmallerThanChild(parentIdx) {
		lcIdx, rcIdx := leftChildIdx(parentIdx), rightChildIdx(parentIdx)
		if heap.data[lcIdx] > heap.data[rcIdx] {
			heap.swapValues(parentIdx, lcIdx)
			parentIdx = lcIdx
		} else {
			heap.swapValues(parentIdx, rcIdx)
			parentIdx = rcIdx
		}
	}
}

func (heap *MaxHeap) swapValues(idx1, idx2 int) {
	heap.data[idx1], heap.data[idx2] = heap.data[idx2], heap.data[idx1]
}

func (heap *MaxHeap) parentSmallerThanChild(parentIdx int) bool {
	parent := heap.data[parentIdx]
	lcIdx, rcIdx := leftChildIdx(parentIdx), rightChildIdx(parentIdx)
	if heap.childExists(lcIdx) && parent < heap.data[lcIdx] {
		return true
	}
	if heap.childExists(rcIdx) && parent < heap.data[rcIdx] {
		return true
	}
	return false
}

func (heap *MaxHeap) childExists(idx int) bool {
	return idx <= len(heap.data)-1
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
