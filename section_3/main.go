package main

import "fmt"

func main() {
	fmt.Println("bar")
	foo()
	for i := 0; i < 100; i++ {
		foo()
	}
}
func foo() {
	fmt.Println("foo")
}
