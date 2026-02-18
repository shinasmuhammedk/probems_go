package functionany

import "fmt"

type Number interface {
	any
}

func Change[T any](a T) {
	fmt.Println(a)
}