package main

import "fmt"

// Stack represents a LIFO (Last-In, First-Out) data structure.
type Stack struct {
	items []int
}

// Push adds an item to the top of the stack.
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the item at the top of the stack.
// Returns an error if the stack is empty.
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	top := len(s.items) - 1
	item := s.items[top]
	s.items = s.items[:top]
	return item, nil
}

// Peek returns the item at the top of the stack without removing it.
// Returns an error if the stack is empty.
func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty returns true if the stack contains no items.
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack.
func (s *Stack) Size() int {
	return len(s.items)
}

func main() {
	s := &Stack{}

	fmt.Println("=== Stack (LIFO) Demo ===")

	// Push items
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Printf("Pushed 1, 2, 3 | Size: %d\n", s.Size())

	// Peek at top
	top, err := s.Peek()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Peek (top): %d\n", top)

	// Pop items
	item, err := s.Pop()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Popped: %d | Size: %d\n", item, s.Size())

	item, err = s.Pop()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Popped: %d | Size: %d\n", item, s.Size())

	item, err = s.Pop()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Popped: %d | Size: %d\n", item, s.Size())

	// Try to pop from empty stack
	_, err = s.Pop()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
