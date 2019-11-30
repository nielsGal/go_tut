package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even, odd)
	go receive(even, odd, fanin)

	for v := range fanin {
		fmt.Printf("formatted value %v", v)
	}
}

func send(even, odd chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}
func receive(even, odd <-chan int, fanin chan<- int) {
	wg.Add(2)
	go func() {
		for v := range even {
			fanin <- v + 10
		}
		wg.Done()
	}()
	go func() {
		for v := range odd {
			fanin <- v - 10
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanin)
}
