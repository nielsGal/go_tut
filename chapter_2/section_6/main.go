package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	x := 2
	fmt.Printf("%d %b\n", x, x)
	x = x << 1
	fmt.Printf("%d %b\n", x, x)

	fmt.Printf("%d %b\n", kb, kb)
	fmt.Printf("%d %b\n", mb, mb)
	fmt.Printf("%d %b\n", gb, gb)
}
