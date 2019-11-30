package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	go send(even, odd, quit)
	receive(even, odd, quit)
}

func send(even, odd, quit chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)

	quit <- 42
	close(quit)
}
func receive(even, odd, quit <-chan int) {
	for {
		select {
		case v, ok := <-even:
			if ok {
				fmt.Printf("even number %v", v)
			} else {
				fmt.Println("even closed")
			}
		case v, ok := <-odd:
			if ok {
				fmt.Printf("odd number %v", v)
			} else {
				fmt.Println("odd closed")
			}
		case v := <-quit:
			fmt.Printf("quitting %v", v)

			return
		}
	}
}
