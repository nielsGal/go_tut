package main

import "fmt"

func main() {
	fmt.Println("bar")
	foo()
}
func foo() {
	fmt.Println("foo")
}
