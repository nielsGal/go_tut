package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 42
	}()

	return c
}
func receive(c1, c2 <-chan int) {
	for {
		select {
		case v, ok := <-c1:
			if ok {
				fmt.Println(v)
			} else {
				return
			}
		case q := <-c2:
			fmt.Println(q)
			return
		}

	}
}
