package main

import "fmt"

var a int = 42

type hotdog int

var b hotdog

type specialInt = int

var c specialInt

func main() {
	b = 52 //assigning b can just be a normal int
	c = 43
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	// cannot do a = b
	a = c // this is clearly possible
	fmt.Println(a)
	a = int(b) // we can do this because hotdog just is int
	fmt.Println(a)
}
