package main

import "fmt"

func main() {
	x := 42
	fmt.Printf("%d %b %#x\n", x, x, x)
	x = x << 2
	fmt.Printf("%d %b %#x", x, x, x)
}
