package stackusingQueue

import "fmt"

type Stack struct {
	Queue []int
}

func (s *Stack) Push(x int) {
	s.Queue = append(s.Queue, x)

	size := len(s.Queue)
	for i := 0; i < size-1; i++ {
		front := s.Queue[0]
		s.Queue = s.Queue[1:]
		s.Queue = append(s.Queue, front)
	}
}

func (s *Stack) Pop() int {
	if len(s.Queue) == 0 {
		fmt.Println("Nothing to Pop")
		return -1
	}
    
    top := s.Queue[0]
    s.Queue = s.Queue[1:]
    return top
}