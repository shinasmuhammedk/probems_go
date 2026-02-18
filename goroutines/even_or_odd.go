package goroutines

import "fmt"

func EvenOrOdd() {

	chOdd := make(chan bool)
	chEven := make(chan bool)

	go func() {
		for i := 1; i < 10; i += 2 {
			<-chOdd
			fmt.Println("Odd",i)
            chEven <-true
		}
	}()
    
    go func ()  {
        for i:=0;i<=10;i+=2{
            <-chEven
            fmt.Println("Even",i)
            chOdd <- true
        }
    }()
    
    chEven <- true
    select{}
}