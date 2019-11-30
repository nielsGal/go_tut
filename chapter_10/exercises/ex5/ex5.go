package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2)
	c <- 42
	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
