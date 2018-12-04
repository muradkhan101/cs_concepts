package main

// QueueWithStacks is a queue implemented with stacks
type QueueWithStacks struct {
	main ArrayStack
	sub  ArrayStack
}

// Enqueue adds item to queue
func (q *QueueWithStacks) Enqueue(item int) {
	q.main.Push(item)
}

// Dequeue removes item from queue
func (q *QueueWithStacks) Dequeue() int {
	if q.sub.Size != 0 {
		return q.sub.Pop()
	}

	for q.main.Size != 0 {
		q.sub.Push(q.main.Pop())
	}

	return q.sub.Pop()
}
