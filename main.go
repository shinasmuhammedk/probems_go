package main

import (
	validanagram "badge/validAnagram"
	validparanthesis "badge/validParanthesis"

	queueusingstack "badge/queueUsingStack"
	stackusingqueue "badge/stackUsingQueue"
	"fmt"
	"badge/goroutines"
	newandmake "badge/newAndMake"
)

func main() {
	goroutines.EvenOrOdd()
    
    goroutines.CatsAndDogs()
    
    newandmake.NewAndMake()
    
    var s stackusingqueue.Stack
    s.Push(100)
    s.Push(295)
    s.Push(182)
    fmt.Println(s.Pop())
    fmt.Println(s.Queue)
    
    var q queueusingstack.Queue
    q.Enqueue(1000)
    q.Enqueue(5)
    q.Enqueue(84)
    fmt.Println(q.Dequeue())
    
    fmt.Println(validparanthesis.Isvalid("{([({[[]]})])}"))
    
    fmt.Println(validanagram.IsAnagram("shinase", "sawnihs"))
}