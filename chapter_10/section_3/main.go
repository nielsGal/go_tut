package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)
	go foo(c)
	bar(c)
	fmt.Println("about to exit")
}
func foo(c chan<- int) { // a channel can be converted from bidirectional to send only
	c <- 42

}

func bar(c <-chan int) { // a channel can be converted from bidirectional to receive only
	fmt.Println(<-c)
}
