package main

import "fmt"

func main() {
	c := make(chan int) //this blocks has to be wrapped insides some GO routine
	go func() {
		c <- 42
	}()
	fmt.Println("the number", <-c)
	c2 := make(chan int, 2) //this doesn't block until the buffer is full
	c2 <- 43
	fmt.Println("the other number", <-c2)
}
