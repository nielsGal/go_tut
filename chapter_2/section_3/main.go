package main

import "fmt"

func main() {
	x := 42
	y := 42.333
	fmt.Printf("%T %T %v %v\n", x, y, x, y)
	fmt.Println(float64(x) == y)
}
