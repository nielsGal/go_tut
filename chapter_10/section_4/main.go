package main

import "fmt"

func main() {
	c := make(chan int)
	go foo(c)
	bar(c)

}
func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
func bar(c <-chan int) {
	for v := range c {
		fmt.Printf("receive %v ", v)
	}
}
