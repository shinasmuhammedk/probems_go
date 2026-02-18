package goroutines

import "fmt"

func CatsAndDogs() {
	chCats := make(chan bool)
	chDogs := make(chan bool)

	go func() {
		for i := 0; i <= 10; i++ {
			<-chCats
			fmt.Println("Cats")
            chDogs <- true
		}
	}()
    
    go func (){
        for i:=0;i<=10;i++{
            <-chDogs
            fmt.Println("Dogs")
            chCats <- true
        }
    }()
    
    chCats <- true
    select{}
}