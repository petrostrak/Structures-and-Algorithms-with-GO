/*
A heap is a data structure that is based on the heap property. The heap data structure is
used in selection, graph, and k-way merge algorithms. Operations such as finding,
merging, insertion, key changes, and deleting are performed on heaps. Heaps are part of
the container/heap package in Go. According to the heap order (maximum heap)
property, the value stored at each node is greater than or equal to its children.
If the order is descending, it is referred to as a maximum heap; otherwise, it's a minimum
heap. It is not a sorted data structure, but partially ordered. The following example
shows how to use the container/heap package to create a heap data structure
*/

package main

import (
	"container/heap"
	"fmt"
)

// IntegerHeap type
type IntegerHeap []int

// Length method of type IntegerHeap
func (iheap IntegerHeap) Len() int {
	return len(iheap)
}

// Checks
func (iheap IntegerHeap) Less(i, j int) bool {
	return iheap[i] < iheap[j]
}

// Swaps
func (iheap IntegerHeap) Swap(i, j int) {
	iheap[i], iheap[j] = iheap[j], iheap[i]
}

// Push method of type IntegerHeap
func (iheap *IntegerHeap) Push(heapintf interface{}) {
	*iheap = append(*iheap, heapintf.(int))
}

// Pop method of type IntegerHeap
func (iheap *IntegerHeap) Pop() interface{} {
	var n int
	var x1 int
	var previous IntegerHeap = *iheap
	n = len(previous)
	x1 = previous[n-1]
	*iheap = previous[0 : n-1]
	return x1
}

func main() {
	var intHeap *IntegerHeap = &IntegerHeap{-2, 0, 2, 4, 6}
	heap.Init(intHeap)
	heap.Push(intHeap, -4)
	fmt.Printf("minimum: %d\n", (*intHeap)[0])

	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}
}
