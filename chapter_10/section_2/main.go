package main

import "fmt"

func main() {
	c := make(chan int, 2)

	fmt.Printf("%T \n", c)
	dic := make(chan<- int, 2)

	fmt.Printf("%T \n", dic)

	dic2 := make(<-chan int, 2)

	fmt.Printf("%T \n", dic2)
}
