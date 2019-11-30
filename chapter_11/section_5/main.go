package main

import "fmt"

func main() {
	defer foo()
	go func() {
		defer foo()
		panic("bla")
	}()
}

func foo() {
	r := recover()
	fmt.Println(r)
}
