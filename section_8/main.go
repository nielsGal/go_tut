package main

import "fmt"

var a int = 42

type hotdog int

var b hotdog

type specialInt = int

var c specialInt

func main() {
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

}
