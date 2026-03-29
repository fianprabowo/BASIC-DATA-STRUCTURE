# Stack

A **Stack** is a linear data structure that follows the **LIFO** (Last-In, First-Out) principle. The last element added is the first one to be removed — just like a stack of plates.

## Operations

| Operation | Description                                      | Time Complexity |
|-----------|--------------------------------------------------|-----------------|
| `Push`    | Add an item to the top of the stack              | O(1)            |
| `Pop`     | Remove and return the item at the top            | O(1)            |
| `Peek`    | Return the top item without removing it          | O(1)            |
| `IsEmpty` | Check whether the stack is empty                 | O(1)            |
| `Size`    | Return the number of items in the stack          | O(1)            |

## Usage

```go
s := &Stack{}

s.Push(1)
s.Push(2)
s.Push(3)

top, _ := s.Peek() // 3
item, _ := s.Pop() // 3
item, _  = s.Pop() // 2
```

## Run

```bash
cd stack
go run main.go
```

## Example Output

```
=== Stack (LIFO) Demo ===
Pushed 1, 2, 3 | Size: 3
Peek (top): 3
Popped: 3 | Size: 2
Popped: 2 | Size: 1
Popped: 1 | Size: 0
Error: stack is empty
```

## Real-World Use Cases

- **Function call stack** – managing program execution frames
- **Undo/Redo** – text editors and applications
- **Depth-First Search (DFS)** – graph/tree traversal
- **Expression evaluation** – parsing and evaluating mathematical expressions
- **Backtracking** – solving mazes, puzzles
