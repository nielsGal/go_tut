package main

import "fmt"

func main() {
	c := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go func(x int) {
			for j := -0; j < 10; j++ {
				c <- x * j
			}
		}(i)
	}
	for i := 0; i < 100; i++ {
		fmt.Println(<-c)
	}
}
