package goroutines

import (
	"fmt"
	"time"
)

func ABC() {
	go func() {
		fmt.Println("A")
	}()

	go func() {
		fmt.Println("B")
	}()
    
    go func(){
        fmt.Println("C")
    }()
    
    time.Sleep(time.Second)
}