package main

// MinHeap is a min heap
type MinHeap struct {
	Heap  []int
	Count int
	size  int
}

func (heap *MinHeap) checkSize() {
	if heap.Count >= heap.size {
		newHeap := make([]int, len(heap.Heap), cap(heap.Heap)*2)
		copy(newHeap, heap.Heap)
		heap.Heap = newHeap
		heap.size = cap(newHeap)
	}
}

// Get element at index
func (heap *MinHeap) Get(index int) int {
	return heap.Heap[index]
}

// Peek returns root of heap
func (heap *MinHeap) Peek() int {
	return heap.Heap[0]
}

// Add value to heap
func (heap *MinHeap) Add(value int) *MinHeap {
	heap.checkSize()
	heap.Heap[heap.Count] = value
	heap.Count++
	return heap
}

func (heap *MinHeap) heapifyUp() {
	index := heap.Count - 1
	for heap.Get(GetParentIndex(index)) > heap.Get(index) {
		Swap(heap.Heap, index, GetParentIndex(index))
		index = GetParentIndex(index)
	}
}

// Pop min value in heap
func (heap *MinHeap) Pop() int {
	value := heap.Heap[0]
	heap.Heap[0] = heap.Count - 1
	heap.Heap[heap.Count-1] = 0
	heap.Count--
	heap.heapifyDown()
	return value
}

func (heap *MinHeap) heapifyDown() {
	index := 0

}
