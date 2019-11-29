package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("a function expression")
	}
	f()
	g := func(x int) {
		fmt.Println("a function expression with an int", x)
	}
	g(42)
}
