package main

import "fmt"

func main() {
	foo()
}

func foo() {
	f := bar()
	fmt.Println("a special function", f())
	fmt.Println("another special function", bar()())
}

func bar() func() int {
	return func() int {
		return 451
	}
}
