# Queue

A **Queue** is a linear data structure that follows the **FIFO** (First-In, First-Out) principle. The first element added is the first one to be removed — just like a line of people waiting.

## Operations

| Operation   | Description                                      | Time Complexity |
|-------------|--------------------------------------------------|-----------------|
| `Enqueue`   | Add an item to the back of the queue             | O(1)            |
| `Dequeue`   | Remove and return the item at the front          | O(n)*           |
| `Peek`      | Return the front item without removing it        | O(1)            |
| `IsEmpty`   | Check whether the queue is empty                 | O(1)            |
| `Size`      | Return the number of items in the queue          | O(1)            |

> \* The slice-based `Dequeue` is O(n) due to element shifting. For high-throughput scenarios, consider a linked-list or circular-buffer implementation to achieve O(1) dequeue.


```go
q := &Queue{}

q.Enqueue(1)
q.Enqueue(2)
q.Enqueue(3)

front, _ := q.Peek()    // 1
item, _  := q.Dequeue() // 1
item, _   = q.Dequeue() // 2
```

## Run

```bash
cd queue
go run main.go
```

## Example Output

```
=== Queue (FIFO) Demo ===
Enqueued 1, 2, 3 | Size: 3
Peek (front): 1
Dequeued: 1 | Size: 2
Dequeued: 2 | Size: 1
Dequeued: 3 | Size: 0
Error: queue is empty
```

## Real-World Use Cases

- **Task scheduling** – OS process queues
- **Breadth-First Search (BFS)** – graph/tree traversal
- **Print spooling** – managing print jobs
- **Message queues** – async communication between services
