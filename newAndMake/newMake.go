package newandmake

import "fmt"

func NewAndMake() {
	ptr := new(int)
    *ptr = 1000
	// fmt.Println(&ptr)
	fmt.Println(*ptr)
    change(ptr)
    fmt.Println(*ptr)
}

func change(ptr *int){
    *ptr = 99
}