package queueusingstack

import "fmt"

type Queue struct {
	stack1 []int
	stack2 []int
}

func (q *Queue) Enqueue(x int) {
	q.stack1 = append(q.stack1, x)
}

func (q *Queue) Dequeue() int {
	if len(q.stack2) == 0 {
		for len(q.stack1) > 0 {
			n := len(q.stack1) - 1
			val := q.stack1[n]
			q.stack1 = q.stack1[:n]
			q.stack2 = append(q.stack2, val)
		}
	}
	if len(q.stack2) == 0 {
		fmt.Println("Queue is empty")
        return -1
	}
    
    n := len(q.stack2)-1
    val := q.stack2[n]
    q.stack2 = q.stack2[:n]
    
    return val
}