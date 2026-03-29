package main

import "fmt"

// Queue represents a FIFO (First-In, First-Out) data structure.
type Queue struct {
	items []int
}

// Enqueue adds an item to the back of the queue.
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the item at the front of the queue.
// Returns an error if the queue is empty.
func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

// Peek returns the item at the front of the queue without removing it.
// Returns an error if the queue is empty.
func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}
	return q.items[0], nil
}

// IsEmpty returns true if the queue contains no items.
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of items in the queue.
func (q *Queue) Size() int {
	return len(q.items)
}

func main() {
	q := &Queue{}

	fmt.Println("=== Queue (FIFO) Demo ===")

	// Enqueue items
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Printf("Enqueued 1, 2, 3 | Size: %d\n", q.Size())

	// Peek at front
	front, err := q.Peek()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Peek (front): %d\n", front)

	// Dequeue items
	item, err := q.Dequeue()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Dequeued: %d | Size: %d\n", item, q.Size())

	item, err = q.Dequeue()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Dequeued: %d | Size: %d\n", item, q.Size())

	item, err = q.Dequeue()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Dequeued: %d | Size: %d\n", item, q.Size())

	// Try to dequeue from empty queue
	_, err = q.Dequeue()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
