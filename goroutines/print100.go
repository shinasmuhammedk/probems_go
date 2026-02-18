package goroutines

import "fmt"

func Print() {

	counter := 1
	max := 100

	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})

	go func() {
		for {
			<-ch1
			if counter > max {
				ch2 <- struct{}{}
				return
			}
			fmt.Println(counter)
            counter++
            ch2 <- struct{}{}
		}
	}()
    
    go func (){
        <-ch2
        if counter > max{
            ch3 <- struct{}{}
            return
        }
        fmt.Println(counter)
        counter++
        ch3 <-struct{}{}
    }()
    
    go func(){
        <- ch3
        if counter > max{
            ch1 <- struct{}{}
            return
        }
        fmt.Println(counter)
        counter++
        ch1 <- struct{}{}
    }()
    
    ch1 <- struct{}{}
    select{}
}